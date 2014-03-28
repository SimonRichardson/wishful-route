package route

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func Route(fallback func() Promise, rs []func(x *Request) Option) func(x *Request) Promise {
	return func(x *Request) Promise {
		opt := compact(rs, x)
		return opt.GetOrElse(func() AnyVal {
			return fallback()
		}).(Promise)
	}
}

// This is a partial applicative
// TODO (simon) : we could implement this as a goroutine
func compact(rs []func(x *Request) Option, x *Request) Option {
	if len(rs) == 0 {
		return NewNone()
	}

	// We should trampoline this!
	a := rs[0](x)
	if b, ok := a.(Some); ok {
		return b
	}

	return compact(rs[1:], x)
}
