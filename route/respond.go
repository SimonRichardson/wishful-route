package route

import (
	"strings"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func respond(method string, path string, responder func(req *Request) AnyVal) func(request *Request) Option {
	lower := strings.ToLower(method)
	extract := CompilePath(path)
	return func(request *Request) Option {
		cond := lower == strings.ToLower(request.Method)
		return guard(cond).Chain(
			func(x AnyVal) Monad {
				url := request.URL.String()
				return extract(url).Chain(
					func(params AnyVal) Monad {
						req := request.SetParams(params.(map[string]string))
						return NewSome(responder(req))
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

func Get(path string, responder func(req *Request) AnyVal) func(request *Request) Option {
	return respond("get", path, responder)
}

func Post(path string, responder func(req *Request) AnyVal) func(request *Request) Option {
	return respond("post", path, responder)
}
