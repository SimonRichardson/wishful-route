package route

import (
	"net/http"
	"strings"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func respond(method string, path string, responder func(req *http.Request) AnyVal) func(request *http.Request) Option {
	lower := strings.ToLower(method)
	extract := CompilePath(path)
	return func(request *http.Request) Option {
		cond := lower == strings.ToLower(request.Method)
		return guard(cond).Chain(
			func(x AnyVal) Monad {
				url := request.URL.String()
				return extract(url).Chain(
					func(params AnyVal) Monad {
						return NewSome(responder(request))
					},
				)
			},
		).(Option)
	}
}

func guard(cond bool) Option {
	if cond {
		return NewSome(Empty{})
	} else {
		return NewNone()
	}
}

func Get(path string, responder func(req *http.Request) AnyVal) func(request *http.Request) Option {
	return respond("get", path, responder)
}

func Post(path string, responder func(req *http.Request) AnyVal) func(request *http.Request) Option {
	return respond("post", path, responder)
}
