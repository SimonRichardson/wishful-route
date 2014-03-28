package route

import (
	"fmt"
	"net/http"
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func Test_ReturnsNoneIfMethodDoesNotMatch(t *testing.T) {
	f := func(x string, y string) Option {
		p := fmt.Sprintf("/%s", x)
		r := respond("GET", p, constantStringReturnPromise(y))
		req, _ := http.NewRequest("POST", p, nil)
		return r(NewRequest(req))
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
		r := respond("GET", p0, constantStringReturnPromise(z))
		req, _ := http.NewRequest("POST", p1, nil)
		return r(NewRequest(req))
	}
	g := func(x string, y string, z string) Option {
		return NewNone()
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

// **
func Test_CallsResponderWithRequestAndReturnsWithSome(t *testing.T) {
	f := func(x string, y string) AnyVal {
		p := fmt.Sprintf("/%s", x)
		r := respond("GET", p, constantStringReturnPromise(y))
		req, _ := http.NewRequest("GET", p, nil)
		return extract(r(NewRequest(req)))
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
		r := Get(p, constantStringReturnPromise(y))
		req, _ := http.NewRequest("GET", p, nil)
		return extract(r(NewRequest(req)))
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
		r := Post(p, constantStringReturnPromise(y))
		req, _ := http.NewRequest("POST", p, nil)
		return extract(r(NewRequest(req)))
	}
	g := func(x string, y string) AnyVal {
		return y
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
