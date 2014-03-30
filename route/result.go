package route

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	resultLens Lens = Lens{}.AccessorLens(resultAccessor{})
)

type resultAccessor struct{}

func (s resultAccessor) Get(x AnyVal) AnyVal {
	return x.(Result).Headers
}

func (s resultAccessor) Set(x AnyVal, y AnyVal) AnyVal {
	r := x.(Result)
	return NewResult(r.Body, r.StatusCode, y.(Headers))
}

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

func (r Result) SetHeaders(h Headers) Result {
	return resultLens.Run(r).Set(h).(Result)
}
