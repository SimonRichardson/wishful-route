package json

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful-route/route"
	. "github.com/SimonRichardson/wishful/wishful"
)

const (
	CallbackName string = "callback"
)

type MockJsonpObject struct {
	String string
}

func NewMockJsonpObject(x string) MockJsonpObject {
	return MockJsonpObject{
		String: x,
	}
}

func NewMockJsonpString(x string) string {
	a, _ := json.Marshal(NewMockJsonpObject(x))
	return Output(CallbackName, string(a))
}

func Output(callback string, result string) string {
	return fmt.Sprintf("%s(%s);", callback, result)
}

func Test_OkShouldReturnCorrectResult(t *testing.T) {
	f := func(x string) Result {
		return Ok(CallbackName, NewMockJsonpObject(x)).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return NewJsonpResult(http.StatusOK, NewMockJsonpString(x))
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_NotFoundShouldReturnCorrectResult(t *testing.T) {
	f := func(x string) Result {
		return NotFound(CallbackName, NewMockJsonpObject(x)).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return NewJsonpResult(http.StatusNotFound, NewMockJsonpString(x))
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_InternalServerErrorShouldReturnCorrectResult(t *testing.T) {
	f := func(x string) Result {
		return InternalServerError(CallbackName, NewMockJsonpObject(x)).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return NewJsonpResult(http.StatusInternalServerError, NewMockJsonpString(x))
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
