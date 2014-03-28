package route

import (
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/useful"
)

func Test_CallsFallbackIfEmptyList(t *testing.T) {
	f := func(x string) string {
		fallback := constantAnyValReturnPromise(x)
		res := Route(fallback, make([]func(x *Request) Option, 0, 0))
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
		fallback := constantAnyValReturnPromise(x)
		routes := []func(x *Request) Option{
			func(a *Request) Option {
				return NewSome(Promise{}.Of(y))
			},
		}
		res := Route(fallback, routes)
		return extractPromise(res(newRequstFromString(z))).(string)
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
		fallback := constantAnyValReturnPromise(x)
		routes := []func(x *Request) Option{
			func(a *Request) Option {
				return NewNone()
			},
			func(a *Request) Option {
				return NewSome(Promise{}.Of(y))
			},
		}
		res := Route(fallback, routes)
		return extractPromise(res(newRequstFromString(z))).(string)
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
		fallback := constantAnyValReturnPromise(x)
		routes := []func(x *Request) Option{
			func(a *Request) Option {
				return NewNone()
			},
			func(a *Request) Option {
				return NewNone()
			},
		}
		res := Route(fallback, routes)
		return extractPromise(res(newRequstFromString(z))).(string)
	}
	g := func(x string, y string, z string) string {
		return x
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
