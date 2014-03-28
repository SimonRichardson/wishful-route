package route

import (
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func extractPromise(x AnyVal) AnyVal {
	promise := x.(Promise)
	return promise.Extract()
}

func constantPromise(x AnyVal) func() Promise {
	return func() Promise {
		return Promise{}.Of(x).(Promise)
	}
}

func Test_CallsFallbackIfEmptyList(t *testing.T) {
	f := func(x string) string {
		fallback := constantPromise(x)
		res := Route(fallback, make([]func(x AnyVal) Option, 0, 0))
		return extractPromise(res(nil)).(string)
	}
	g := func(x string) string {
		return x
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_CallsFirstMatch(t *testing.T) {
	f := func(x string, y string, z string) string {
		fallback := constantPromise(x)
		routes := []func(x AnyVal) Option{
			func(a AnyVal) Option {
				return NewSome(Promise{}.Of(y))
			},
		}
		res := Route(fallback, routes)
		return extractPromise(res(z)).(string)
	}
	g := func(x string, y string, z string) string {
		return y
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_CallsFirstMatchWhenMultiple(t *testing.T) {
	f := func(x string, y string, z string) string {
		fallback := constantPromise(x)
		routes := []func(x AnyVal) Option{
			func(a AnyVal) Option {
				return NewNone()
			},
			func(a AnyVal) Option {
				return NewSome(Promise{}.Of(y))
			},
		}
		res := Route(fallback, routes)
		return extractPromise(res(z)).(string)
	}
	g := func(x string, y string, z string) string {
		return y
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_CallsFallbackIfNoMatch(t *testing.T) {
	f := func(x string, y string, z string) string {
		fallback := constantPromise(x)
		routes := []func(x AnyVal) Option{
			func(a AnyVal) Option {
				return NewNone()
			},
			func(a AnyVal) Option {
				return NewNone()
			},
		}
		res := Route(fallback, routes)
		return extractPromise(res(z)).(string)
	}
	g := func(x string, y string, z string) string {
		return x
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
