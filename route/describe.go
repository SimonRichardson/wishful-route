package route

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func Describe(d string, f func(x *Request) Option) func(x *Request) Option {
	return func(x *Request) Option {
		return f(x).Map(func(a AnyVal) AnyVal {
			return a.(Promise).Map(func(a AnyVal) AnyVal {
				result := a.(Result)
				headers := result.Headers
				return result.SetHeaders(headers.Add(NewTuple2("x-describe", d)))
			}).(Promise)
		}).(Option)
	}
}
