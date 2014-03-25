package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

type Partial struct {
	Value string `json:"value"`
}

func NewPartial(x string) *Partial {
	return &Partial{Value: x}
}

func (x *Partial) Show() (string, string) {
	a, _ := json.Marshal(x)
	return string(a), fmt.Sprintf("%d", len(a))
}

func values(x string) (string, string) {
	v := url.Values{}
	v.Set("xx", x)
	e := v.Encode()
	return e, fmt.Sprintf("%d", len(e))
}

func request(body string, length string) *http.Request {
	reader := strings.NewReader(body)
	closer := ioutil.NopCloser(reader)
	req, _ := http.NewRequest("get", "/", closer)
	req.Header.Set("Content-Length", length)
	return req
}

func Test_RawShouldReturnString(t *testing.T) {
	f := func(x string) string {
		promise := Raw(request(x, fmt.Sprintf("%d", len(x)))).Run.(Promise)
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

func Test_JsonShouldReturnPartial(t *testing.T) {
	f := func(x string) *Partial {
		var partial Partial
		promise := Json(&partial, request(NewPartial(x).Show())).Run.(Promise)
		return promise.Fork(func(x AnyVal) AnyVal {
			return x.(Either).Fold(Identity, Identity)
		}).(*Partial)
	}
	g := func(x string) *Partial {
		return NewPartial(x)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_JsonShouldReturnErrorIfInvalidBody(t *testing.T) {
	f := func(x string) bool {
		a, b := NewPartial(x).Show()

		var partial Partial
		promise := Json(&partial, request(a[0:len(a)-2], b)).Run.(Promise)
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

func Test_JsonShouldReturnErrorIfInvalidBodyWhichShouldThrowParseError(t *testing.T) {
	f := func(x string) bool {
		a := `{"asd`
		b := fmt.Sprintf("%d", len(a))

		var partial Partial
		promise := Json(&partial, request(a, b)).Run.(Promise)
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

func Test_QueryShouldReturnValues(t *testing.T) {
	f := func(x string) url.Values {
		promise := Query(request(values(x))).Run.(Promise)
		return promise.Fork(func(x AnyVal) AnyVal {
			return x.(Either).Fold(Identity, Identity)
		}).(url.Values)
	}
	g := func(x string) url.Values {
		v := url.Values{}
		v.Set("xx", x)
		return v
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_QueryShouldReturnErrorIfInvalidBody(t *testing.T) {
	f := func(x string) bool {
		a, b := values(x)

		promise := Query(request(a[0:len(a)-2], b)).Run.(Promise)
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

func Test_QueryShouldReturnErrorIfInvalidBodyWhichShouldThrowParseError(t *testing.T) {
	f := func(x string) bool {
		a := "%as=hello"
		b := fmt.Sprintf("%d", len(a))

		promise := Query(request(a, b)).Run.(Promise)
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
