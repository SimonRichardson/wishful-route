package route

import (
	"fmt"
	"net/url"
	"regexp"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

const (
	ident string = "[a-z$_][a-z0-9$_]*"
	sigil string = ":"
)

var (
	param string = fmt.Sprintf("%s(%s)", sigil, ident)
)

type params struct {
	RegExp *regexp.Regexp
	Names  []string
}

func flattenIdents(names [][]string) []string {
	n := len(names)
	res := make([]string, n, n)
	for k, v := range names {
		res[k] = v[1]
	}
	return res
}

func flattenNames(names [][]string) []string {
	if len(names) > 0 {
		return names[0][1:]
	}
	return make([]string, 0, 0)
}

func extractIdents(path string) Option {
	r, e := regexp.Compile(param)
	if e != nil {
		return NewNone()
	}
	return NewSome(r.ReplaceAllString(path, "([^\\/]+)"))
}

func extractIdentNames(path string) Option {
	r, e := regexp.Compile(param)
	if e != nil {
		return NewNone()
	}
	a := r.FindAllStringSubmatch(path, -1)
	if len(a) > 0 {
		return NewSome(flattenIdents(a))
	}
	return NewNone()
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

func getParams(path string) Monad {
	paramReg := extractIdents(path)
	return paramReg.Chain(
		func(x AnyVal) Monad {
			return compileReg(path, x.(string))
		},
	).Chain(
		func(x AnyVal) Monad {
			expr := x.(*regexp.Regexp)
			return extractIdentNames(path).Map(
				func(x AnyVal) AnyVal {
					return params{
						RegExp: expr,
						Names:  x.([]string),
					}
				},
			).(Option).OrElse(getDefaultNames(expr))
		},
	).(Option)
}

func getDefaultNames(expr *regexp.Regexp) Option {
	return NewSome(params{
		RegExp: expr,
		Names:  make([]string, 0, 0),
	})
}

func CompilePath(path string) func(url string) Option {
	reg := getParams(path)
	return func(raw string) Option {
		return reg.Chain(func(x AnyVal) Monad {
			p := x.(params)

			u, err := url.Parse(raw)
			if err != nil {
				return NewNone()
			}

			exp := p.RegExp
			that := exp.FindAllStringSubmatch(u.Path, -1)

			if len(that) < 1 {
				return NewNone()
			}

			res := make(map[string]string)
			for k, v := range flattenNames(that) {
				res[p.Names[k]] = v
			}
			return NewSome(res)
		}).(Option)
	}
}

func CompileSimplePath(path string) func(url string) Option {
	return func(raw string) Option {
		u, err := url.Parse(raw)
		if err != nil || path != u.Path {
			return NewNone()
		}

		return NewSome(make(map[string]string))
	}
}
