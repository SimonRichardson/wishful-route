package route

import (
	"fmt"
	"net/http"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func extract(x Option) AnyVal {
	return x.Fold(
		func(y AnyVal) AnyVal {
			if a, ok := y.(Promise); ok {
				return a.Extract()
			}
			return y
		},
		func() AnyVal {
			panic(fmt.Errorf("Failed if called"))
			return Empty{}
		},
	)
}

func extractPromise(x AnyVal) AnyVal {
	promise := x.(Promise)
	return promise.Extract()
}

func constantStringReturnPromise(x string) func(req *Request) Promise {
	return func(req *Request) Promise {
		return Promise{}.Of(x).(Promise)
	}
}

func constantAnyValReturnPromise(x AnyVal) func() Promise {
	return func() Promise {
		return Promise{}.Of(x).(Promise)
	}
}

func newRequstFromString(s string) *Request {
	req, _ := http.NewRequest("GET", s, nil)
	return NewRequest(req)
}
