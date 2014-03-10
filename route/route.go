package route

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

type RouteCallback func(x AnyVal) Option

func Route(fallback AnyVal, rs []RouteCallback) func(x AnyVal) AnyVal {
	return func(x AnyVal) AnyVal {
		opt := compact(rs, x)
		return opt.GetOrElse(fallback)
	}
}

// This is a partial applicative
// TODO (simon) : we could implement this as a goroutine
func compact(rs []RouteCallback, x AnyVal) Option {
	if len(rs) == 0 {
		return NewNone()
	}

	a := rs[0](x)
	if b, ok := a.(Some); ok {
		return b
	}

	return compact(rs[1:], x)
}
