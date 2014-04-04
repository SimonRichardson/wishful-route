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

func Ok(req *Request, body AnyVal) Promise {
	switch getFormat(req.Request) {
	case XmlFormat:
		panic(errors.New("Missing implementation"))
	case TxtFormat:
		return plainResult.Ok(body.(string))
	case JsonFormat:
		return jsonResult.Ok(body)
	default:
		return htmlResult.Ok(body.(string))
	}
}

func NotFound(req *Request, body AnyVal) Promise {
	switch getFormat(req.Request) {
	case XmlFormat:
		panic(errors.New("Missing implementation"))
	case TxtFormat:
		return plainResult.NotFound(body.(string))
	case JsonFormat:
		return jsonResult.NotFound(body)
	default:
		return htmlResult.NotFound(body.(string))
	}
}

func InternalServerError(req *Request, body AnyVal) Promise {
	switch getFormat(req.Request) {
	case XmlFormat:
		panic(errors.New("Missing implementation"))
	case TxtFormat:
		return plainResult.InternalServerError(body.(string))
	case JsonFormat:
		return jsonResult.InternalServerError(body)
	default:
		return htmlResult.InternalServerError(body.(string))
	}
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
	default:
		return HtmlFormat
	}
}
