package route

import (
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func Test_CallsFallbackIfEmptyList(t *testing.T) {
	f := func(x string) string {
		fallback := ConstantNoArgs(x)
		res := Route(fallback, make([]func(x AnyVal) Option, 0, 0))
		return res(nil).(string)
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
		fallback := ConstantNoArgs(x)
		routes := []func(x AnyVal) Option{
			func(a AnyVal) Option {
				return NewSome(y)
			},
		}
		res := Route(fallback, routes)
		return res(z).(string)
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
		fallback := ConstantNoArgs(x)
		routes := []func(x AnyVal) Option{
			func(a AnyVal) Option {
				return NewNone()
			},
			func(a AnyVal) Option {
				return NewSome(y)
			},
		}
		res := Route(fallback, routes)
		return res(z).(string)
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
		fallback := ConstantNoArgs(x)
		routes := []func(x AnyVal) Option{
			func(a AnyVal) Option {
				return NewNone()
			},
			func(a AnyVal) Option {
				return NewNone()
			},
		}
		res := Route(fallback, routes)
		return res(z).(string)
	}
	g := func(x string, y string, z string) string {
		return x
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
