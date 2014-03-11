package route

import (
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func handleException(f func(x AnyVal) Either) func(x AnyVal) Either {
	return func(x AnyVal) AnyVal {
		fun := NewFunction(f)
		res, err := fun.Call(x)
		if err != nil {
			return NewLeft(err)
		} else {
			return res
		}
	}
}
