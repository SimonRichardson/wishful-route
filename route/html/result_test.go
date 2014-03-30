package html

import (
	"net/http"
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful-route/route"
	. "github.com/SimonRichardson/wishful/wishful"
)

func Test_OkShouldReturnCorrectResult(t *testing.T) {
	f := func(x string) Result {
		return Ok(x).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return NewHtmlResult(http.StatusOK, x)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_NotFoundShouldReturnCorrectResult(t *testing.T) {
	f := func(x string) Result {
		return NotFound(x).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return NewHtmlResult(http.StatusNotFound, x)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_InternalServerErrorShouldReturnCorrectResult(t *testing.T) {
	f := func(x string) Result {
		return InternalServerError(x).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return NewHtmlResult(http.StatusInternalServerError, x)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
