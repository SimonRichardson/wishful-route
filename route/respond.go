package route

import (
	"strings"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	ParamLens Lens = Lens{}.AccessorLens(ParamAccessor{})
)

type ParamAccessor struct{}

func (s ParamAccessor) Get(x AnyVal) AnyVal {
	return x.(*Request).Params
}

func (s ParamAccessor) Set(x AnyVal, y AnyVal) AnyVal {
	a := NewRequest(x.(*Request).Request)
	a.Params = y.(map[string]string)
	return a
}

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
						req := ParamLens.Run(request).Set(params)
						return NewSome(responder(req.(*Request)))
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
