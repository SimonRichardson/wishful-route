package route

import (
	"fmt"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
	"net/url"
	"regexp"
)

const (
	ident string = "[a-z$_][a-z0-9$_]*"
	sigil string = ":"
)

var (
	param string = fmt.Sprintf("%s(%s)", sigil, ident)
)

func extractIdents(path string) Option {
	r, e := regexp.Compile(param)
	if e != nil {
		return NewNone()
	}
	return NewSome(r.ReplaceAllFunc([]byte(path), func(a []byte) []byte {
		return []byte("([^\\/]+)")
	}))
}

func compileReg(path string, reg string) Option {
	if path == "/" {
		return NewSome(regexp.MustCompile("^/$"))
	}

	var a string

	r := regexp.MustCompile("/$")
	if r.Match([]byte(path)) {
		a = fmt.Sprintf("^%s?", reg)
	} else {
		a = fmt.Sprintf("^%s\\/?", reg)
	}

	exp, err := regexp.Compile(a)
	if err != nil {
		return NewNone()
	}
	return NewSome(exp)
}

func CompilePath(path string) func(url string) Option {
	paramReg := extractIdents(path)
	reg := paramReg.Chain(func(x AnyVal) Monad {
		return compileReg(path, x.(string))
	}).(Option)
	return func(raw string) Option {
		return reg.Chain(func(x AnyVal) Monad {
			u, err := url.Parse(raw)
			if err != nil {
				return NewNone()
			}
			// Retrieve the path
			pathName := []byte(u.Path)

			exp := x.(regexp.Regexp)
			that := exp.FindAll(pathName, -1)
			if that != nil {
				return NewNone()
			}

			result := make([]string, 0, 0)
			for k, v := range that {
				result = append(result, string(v))
			}

			return NewSome(result)
		}).(Option)
	}
}
