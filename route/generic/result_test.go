package generic

import (
	"net/http"
	"testing"
	"testing/quick"
	. "github.com/SimonRichardson/wishful-route/route"
	"github.com/SimonRichardson/wishful-route/route/html"
	. "github.com/SimonRichardson/wishful/wishful"
)

func Test_OkShouldReturnCorrectResultForHtml(t *testing.T) {
	f := func(x string) Result {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Accept", "html")
		return Ok(NewRequest(req), x).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return html.NewHtmlResult(http.StatusOK, x)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_NotFoundShouldReturnCorrectResultForHtml(t *testing.T) {
	f := func(x string) Result {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Accept", "html")
		return NotFound(NewRequest(req), x).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return html.NewHtmlResult(http.StatusNotFound, x)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}

func Test_InternalServerErrorShouldReturnCorrectResultForHtml(t *testing.T) {
	f := func(x string) Result {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("Accept", "html")
		return InternalServerError(NewRequest(req), x).Fork(Identity).(Result)
	}
	g := func(x string) Result {
		return html.NewHtmlResult(http.StatusInternalServerError, x)
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
