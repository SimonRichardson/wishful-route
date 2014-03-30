package html

import (
	"fmt"
	"net/http"
	. "github.com/SimonRichardson/wishful-route/route"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func NewHtmlResult(statusCode int, body string) Result {
	return NewResult(body, statusCode, NewHeaders(map[string]string{
		"Content-Length": fmt.Sprintf("%d", len(body)),
		"Content-Type":   "text/html; charset=utf-8",
	}))
}

func Ok(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewHtmlResult(http.StatusOK, body))
	})
}

func NotFound(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewHtmlResult(http.StatusNotFound, body))
	})
}

func InternalServerError(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewHtmlResult(http.StatusInternalServerError, body))
	})
}

func UseTemplate(x string) Promise {
	return Promise{}.Of(x).(Promise)
}
