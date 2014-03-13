package route

import (
	"errors"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func handleException(f func(x AnyVal) Either) func(x AnyVal) Either {
	return func(x AnyVal) Either {
		fun := NewFunction(f)
		res, err := fun.Call(x)
		if err != nil {
			return NewLeft(err)
		} else {
			if e, ok := res.(Either); ok {
				return e
			}
			return NewLeft(errors.New("TypeError: Not type of Either"))
		}
	}
}
