package route

import (
	"fmt"
	"net/http"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

type Headers map[string]string

func NewHeaders(m map[string]string) Headers {
	return Headers(m)
}

func (h Headers) Add(t Tuple2) Headers {
	m := map[string]string(h)
	r := make(map[string]string)
	for k, v := range m {
		r[k] = v
	}
	r[t.Get1().(string)] = t.Get2().(string)
	return Headers(r)
}

type Result struct {
	Body       string
	StatusCode int
	Headers    Headers
}

func NewResult(body string, statusCode int, headers Headers) Result {
	return Result{
		Body:       body,
		StatusCode: statusCode,
		Headers:    headers,
	}
}

func (r Result) Plain(statusCode int, body string) Result {
	return NewResult(body, statusCode, NewHeaders(map[string]string{
		"Content-Length": fmt.Sprintf("%d", len(body)),
		"Content-Type":   "text/plain",
	}))
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
		return resolve(NewResult("", http.StatusFound, NewHeaders(map[string]string{
			"Location": url,
		})))
	})
}
