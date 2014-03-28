package route

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func Route(fallback func() AnyVal, rs []func(x AnyVal) Option) func(x AnyVal) AnyVal {
	return func(x AnyVal) AnyVal {
		opt := compact(rs, x)
		return opt.GetOrElse(fallback())
	}
}

// This is a partial applicative
// TODO (simon) : we could implement this as a goroutine
func compact(rs []func(x AnyVal) Option, x AnyVal) Option {
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
