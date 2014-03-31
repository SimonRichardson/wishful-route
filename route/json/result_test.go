package json

import (
	"encoding/json"
	"net/http"
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful-route/route"
	. "github.com/SimonRichardson/wishful/wishful"
)

type MockJsonObject struct {
	String string
}

func NewMockJsonObject(x string) MockJsonObject {
	return MockJsonObject{
		String: x,
	}
}

func NewMockJsonString(x string) string {
	a, _ := json.Marshal(NewMockJsonObject(x))
	return string(a)
}

func Test_OkShouldReturnCorrectResult(t *testing.T) {
	f := func(x string) Result {
		return Ok(NewMockJsonObject(x)).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return NewJsonResult(http.StatusOK, NewMockJsonString(x))
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_NotFoundShouldReturnCorrectResult(t *testing.T) {
	f := func(x string) Result {
		return NotFound(NewMockJsonObject(x)).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return NewJsonResult(http.StatusNotFound, NewMockJsonString(x))
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_InternalServerErrorShouldReturnCorrectResult(t *testing.T) {
	f := func(x string) Result {
		return InternalServerError(NewMockJsonObject(x)).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return NewJsonResult(http.StatusInternalServerError, NewMockJsonString(x))
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
