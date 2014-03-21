package route

import (
	"fmt"
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func Test_CompilerShouldReturnSomeOnMatchingPath(t *testing.T) {
	f := func(x string) AnyVal {
		p := fmt.Sprintf("/%s", x)
		return extract(CompilePath(p)(p))
	}
	g := func(x string) AnyVal {
		return make(map[string]string)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_CompilerShouldReturnNoneOnANonMatchingPath(t *testing.T) {
	f := func(x string, y string) Option {
		p0 := fmt.Sprintf("/%s", x)
		p1 := fmt.Sprintf("/nope/%s", y)
		return CompilePath(p0)(p1)
	}
	g := func(x string, y string) Option {
		return NewNone()
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_CompilerShouldReturnSomeOnMatchingPathWithTrailingSlash(t *testing.T) {
	f := func(x string) AnyVal {
		p0 := fmt.Sprintf("/%s/", x)
		p1 := fmt.Sprintf("/%s", x)
		u := extract(CompilePath(p0)(p1))
		return u
	}
	g := func(x string) AnyVal {
		return make(map[string]string)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_CompilerShouldReturnSomeOnMatchingPathWithPrefix(t *testing.T) {
	f := func(x string) AnyVal {
		p0 := fmt.Sprintf("/x%s/", x)
		p1 := fmt.Sprintf("/x%s/trail", x)
		return extract(CompilePath(p0)(p1))
	}
	g := func(x string) AnyVal {
		return make(map[string]string)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_CompilerShouldReturnSomeWithParameter(t *testing.T) {
	f := func(x string) AnyVal {
		p := fmt.Sprintf("/x%s", x)
		u := extract(CompilePath("/:param")(p))
		return u
	}
	g := func(x string) AnyVal {
		p := fmt.Sprintf("x%s", x)
		return map[string]string{
			"param": p,
		}
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
