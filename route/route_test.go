package route

import (
	. "github.com/SimonRichardson/wishful/wishful"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	fallback := Constant("Hello")
	res := Route(fallback, make([]RouteCallback, 0, 0))
	assert.Equal(t, res(nil), "Hello")
}
