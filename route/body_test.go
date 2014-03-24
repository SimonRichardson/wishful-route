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

func request(body string) *http.Request {
	reader := strings.NewReader(body)
	closer := ioutil.NopCloser(reader)
	req, _ := http.NewRequest("get", "/", closer)
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(body)))
	return req
}

func Test_ReadBodyShouldReturnEitherT(t *testing.T) {
	f := func(x string) string {
		promise := ReadBody(request(x)).Run.(Promise)
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
