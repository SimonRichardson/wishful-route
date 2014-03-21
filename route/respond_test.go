package route

import (
	"fmt"
	"net/http"
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func extract(x Option) AnyVal {
	return x.Fold(
		Identity,
		func() AnyVal {
			panic(fmt.Errorf("Failed if called"))
			return Empty{}
		},
	)
}

func constant(x string) func(req *Request) AnyVal {
	return func(req *Request) AnyVal {
		return x
	}
}

func Test_ReturnsNoneIfMethodDoesNotMatch(t *testing.T) {
	f := func(x string, y string) Option {
		p := fmt.Sprintf("/%s", x)
		r := respond("GET", p, constant(y))
		req, _ := http.NewRequest("POST", p, nil)
		return r(req)
	}
	g := func(x string, y string) Option {
		return NewNone()
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ReturnsNoneIfPathDoesNotMatch(t *testing.T) {
	f := func(x string, y string, z string) Option {
		p0 := fmt.Sprintf("/%s", x)
		p1 := fmt.Sprintf("/%s", y)
		r := respond("GET", p0, constant(z))
		req, _ := http.NewRequest("POST", p1, nil)
		return r(req)
	}
	g := func(x string, y string, z string) Option {
		return NewNone()
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_CallsResponderWithRequestAndReturnsWithSome(t *testing.T) {
	f := func(x string, y string) AnyVal {
		p := fmt.Sprintf("/%s", x)
		r := respond("GET", p, constant(y))
		req, _ := http.NewRequest("GET", p, nil)
		return extract(r(req))
	}
	g := func(x string, y string) AnyVal {
		return y
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_CallsGetWithRequestAndReturnsWithSome(t *testing.T) {
	f := func(x string, y string) AnyVal {
		p := fmt.Sprintf("/%s", x)
		r := Get(p, constant(y))
		req, _ := http.NewRequest("GET", p, nil)
		return extract(r(req))
	}
	g := func(x string, y string) AnyVal {
		return y
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_CallsPostWithRequestAndReturnsWithSome(t *testing.T) {
	f := func(x string, y string) AnyVal {
		p := fmt.Sprintf("/%s", x)
		r := Post(p, constant(y))
		req, _ := http.NewRequest("POST", p, nil)
		return extract(r(req))
	}
	g := func(x string, y string) AnyVal {
		return y
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
