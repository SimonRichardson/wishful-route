package plain

import (
	"fmt"
	"net/http"
	. "github.com/SimonRichardson/wishful-route/route"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func NewPlainResult(statusCode int, body string) Result {
	return NewResult(body, statusCode, NewHeaders(map[string]string{
		"Content-Length": fmt.Sprintf("%d", len(body)),
		"Content-Type":   "text/plain",
	}))
}

func Ok(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewPlainResult(http.StatusOK, body))
	})
}

func NotFound(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewPlainResult(http.StatusNotFound, body))
	})
}

func InternalServerError(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewPlainResult(http.StatusInternalServerError, body))
	})
}

func Redirect(url string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewResult("", http.StatusFound, NewHeaders(map[string]string{
			"Location": url,
		})))
	})
}
