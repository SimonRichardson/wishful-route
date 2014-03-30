package route

import (
	"strings"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func compilePaths(path Either) func(url string) Option {
	return path.Fold(
		func(x AnyVal) AnyVal {
			return CompileSimplePath(x.(string))
		},
		func(x AnyVal) AnyVal {
			return CompilePath(x.(string))
		},
	).(func(url string) Option)
}

func respond(method string, path Either, responder func(req *Request) Promise) func(request *Request) Option {
	lower := strings.ToLower(method)
	extract := compilePaths(path)
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

func Get(path Either, responder func(req *Request) Promise) func(request *Request) Option {
	return respond("get", path, responder)
}

func Post(path Either, responder func(req *Request) Promise) func(request *Request) Option {
	return respond("post", path, responder)
}

func Put(path Either, responder func(req *Request) Promise) func(request *Request) Option {
	return respond("put", path, responder)
}

func Patch(path Either, responder func(req *Request) Promise) func(request *Request) Option {
	return respond("patch", path, responder)
}

func Delete(path Either, responder func(req *Request) Promise) func(request *Request) Option {
	return respond("delete", path, responder)
}

func Options(path Either, responder func(req *Request) Promise) func(request *Request) Option {
	return respond("options", path, responder)
}
