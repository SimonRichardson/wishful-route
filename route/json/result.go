package json

import (
	"encoding/json"
	"fmt"
	"net/http"
	. "github.com/SimonRichardson/wishful-route/route"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

type Error struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewJsonResult(statusCode int, body string) Result {
	return NewResult(body, statusCode, NewHeaders(map[string]string{
		"Content-Length": fmt.Sprintf("%d", len(body)),
		"Content-Type":   "application/json; charset=utf-8",
	}))
}

func NewError(e error) Error {
	return Error{
		Title:       "Oops! a server error has occurred, marshalling JSON.",
		Description: e.Error(),
	}
}

func marshal(body AnyVal) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		a, e := json.Marshal(body)
		if e != nil {
			b, _ := json.Marshal(NewError(e))
			return resolve(NewLeft(b))
		}
		return resolve(NewRight(string(a)))
	})
}

func Ok(body AnyVal) Promise {
	return marshal(body).Map(func(x AnyVal) AnyVal {
		return x.(Either).Fold(
			func(a AnyVal) AnyVal {
				return NewJsonResult(http.StatusInternalServerError, a.(string))
			},
			func(a AnyVal) AnyVal {
				return NewJsonResult(http.StatusOK, a.(string))
			},
		)
	}).(Promise)
}

func NotFound(body AnyVal) Promise {
	return marshal(body).Map(func(x AnyVal) AnyVal {
		return x.(Either).Fold(
			func(a AnyVal) AnyVal {
				return NewJsonResult(http.StatusInternalServerError, a.(string))
			},
			func(a AnyVal) AnyVal {
				return NewJsonResult(http.StatusNotFound, a.(string))
			},
		)
	}).(Promise)
}

func InternalServerError(body AnyVal) Promise {
	return marshal(body).Map(func(x AnyVal) AnyVal {
		return x.(Either).Fold(
			func(a AnyVal) AnyVal {
				return NewJsonResult(http.StatusInternalServerError, a.(string))
			},
			func(a AnyVal) AnyVal {
				return NewJsonResult(http.StatusInternalServerError, a.(string))
			},
		)
	}).(Promise)
}
