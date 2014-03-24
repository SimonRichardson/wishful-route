package route

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

func request(body string, length string) *http.Request {
	reader := strings.NewReader(body)
	closer := ioutil.NopCloser(reader)
	req, _ := http.NewRequest("get", "/", closer)
	req.Header.Set("Content-Length", length)
	return req
}

func Test_ReadBodyShouldReturnString(t *testing.T) {
	f := func(x string) string {
		promise := ReadBody(request(x, fmt.Sprintf("%d", len(x)))).Run.(Promise)
		return promise.Fork(func(x AnyVal) AnyVal {
			return x.(Either).Fold(Identity, Identity)
		}).(string)
	}
	g := func(x string) string {
		return x
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ReadBodyShouldReturnErrorForInvalidContentLength(t *testing.T) {
	f := func(x string) bool {
		promise := ReadBody(request(x, "abc")).Run.(Promise)
		return promise.Fork(func(x AnyVal) AnyVal {
			return x.(Either).Fold(Constant(true), Constant(false))
		}).(bool)
	}
	g := func(x string) bool {
		return true
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_ReadBodyShouldReturnErrorForContentLengthMismatch(t *testing.T) {
	f := func(x string) bool {
		promise := ReadBody(request(x, fmt.Sprintf("%d", len(x)+1))).Run.(Promise)
		return promise.Fork(func(x AnyVal) AnyVal {
			return x.(Either).Fold(Constant(true), Constant(false))
		}).(bool)
	}
	g := func(x string) bool {
		return true
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
