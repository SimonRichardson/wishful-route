package route

import (
	"fmt"
	"net/http"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

type Result struct {
	Body       string
	StatusCode int
	Headers    map[string]string
}

func NewResult(body string, statusCode int, headers map[string]string) Result {
	return Result{
		Body:       body,
		StatusCode: statusCode,
		Headers:    headers,
	}
}

func (r Result) Plain(statusCode int, body string) Result {
	return NewResult(body, statusCode, map[string]string{
		"Content-Length": fmt.Sprintf("%d", len(body)),
		"Content-Type":   "text/plain",
	})
}

func Ok(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(Result{}.Plain(http.StatusOK, body))
	})
}

func NotFound(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(Result{}.Plain(http.StatusNotFound, body))
	})
}

func InternalServerError(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(Result{}.Plain(http.StatusInternalServerError, body))
	})
}

func Redirect(url string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewResult("", http.StatusFound, map[string]string{
			"Location": url,
		}))
	})
}
