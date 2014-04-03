package generic

import (
	"errors"
	"net/http"
	"strings"
	. "github.com/SimonRichardson/wishful-route/route"
	htmlResult "github.com/SimonRichardson/wishful-route/route/html"
	jsonResult "github.com/SimonRichardson/wishful-route/route/json"
	plainResult "github.com/SimonRichardson/wishful-route/route/plain"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

type Format string

const (
	HtmlFormat  Format = "html"
	XmlFormat   Format = "xml"
	TxtFormat   Format = "txt"
	JsonFormat  Format = "json"
	JsonpFormat Format = "jsonp"
)

func NewGenericResult(req *Request, statusCode int, body string) Result {
	switch getFormat(req.Request) {
	case XmlFormat:
		panic(errors.New("Missing implementation"))
	case TxtFormat:
		return plainResult.NewPlainResult(statusCode, body)
	case JsonFormat:
		return jsonResult.NewJsonResult(statusCode, body)
	default:
		return htmlResult.NewHtmlResult(statusCode, body)
	}
}

func Ok(req *Request, body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewGenericResult(req, http.StatusOK, body))
	})
}

func NotFound(req *Request, body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewGenericResult(req, http.StatusNotFound, body))
	})
}

func InternalServerError(req *Request, body string) Promise {
	return NewPromise(func(resolve func(x AnyVal) AnyVal) AnyVal {
		return resolve(NewGenericResult(req, http.StatusInternalServerError, body))
	})
}

func getFormat(req *http.Request) Format {
	accept := req.Header.Get("accept")

	switch {
	case strings.Contains(accept, "application/xml"),
		strings.Contains(accept, "text/xml"):
		return XmlFormat
	case strings.Contains(accept, "text/plain"):
		return TxtFormat
	case strings.Contains(accept, "application/json"):
		return JsonFormat
	case strings.Contains(accept, "text/javascript"):
		return JsonpFormat
	}

	return HtmlFormat
}
