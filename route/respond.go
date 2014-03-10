package route

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
	"net/http"
	"strings"
)

func respond(method string, path string, responder func(req Request) AnyVal) Option {
	lower := strings.ToLower(method)
	extract := compilePath(path)
	return func(request http.Request) Option {
		return guard(lower == strings.ToLower(request.Method)).Chain(func(x AnyVal) Monad {
			return extract(request.URL).Chain(func(params AnyVal) Monad {
				req := NewRequest(request)
				return Some(responder(req))
			})
		})
	}
}

func guard(cond bool) Monad {
	if cond {
		return NewSome(Empty{})
	} else {
		return NewNone()
	}
}

func Get(path string, responder func(req Request) AnyVal) Option {
	return respond("get", path, responder)
}

func Post(path string, responder func(req Request) AnyVal) Option {
	return respond("post", path, responder)
}
