package route

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
	"net/http"
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
		Header:     headers,
	}
}

func (r Result) Plain(statusCode int, body string) Result {
	return NewResult(body, statusCode, map[string]string{
		"Content-Length": len(body),
		"Content-Type":   "text/plain",
	})
}

func (r Result) Ok(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(r.Plain(http.StatusOK, body))
	})
}

func (r Result) NotFound(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(r.Plain(http.StatusNotFound, body))
	})
}

func (r Result) InternalServerError(body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(r.Plain(http.StatusInternalServerError, body))
	})
}

func (r Result) Redirect(url string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewResult("", 302, map[string]string{
			"Location": url,
		}))
	})
}
