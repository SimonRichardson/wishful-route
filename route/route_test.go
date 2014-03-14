package route

import (
	"testing"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	fallback := ConstantNoArgs("Hello")
	res := Route(fallback, make([]func(x AnyVal) Option, 0, 0))
	assert.Equal(t, res(nil), "Hello")
}
