package route

import (
	"net/http"
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/wishful"
)

func Test_OkShouldReturnCorrectResult(t *testing.T) {
	f := func(x string) Result {
		return Ok(x).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return Result{}.Plain(http.StatusOK, x)
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
		return Result{}.Plain(http.StatusNotFound, x)
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
		return Result{}.Plain(http.StatusInternalServerError, x)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_RedirectShouldReturnCorrectResult(t *testing.T) {
	f := func(x string) Result {
		return Redirect(x).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return NewResult("", http.StatusFound, NewHeaders(map[string]string{
			"Location": x,
		}))
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
