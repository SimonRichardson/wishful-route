package route

import (
	"net/http"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	requestLens Lens = Lens{}.AccessorLens(requestAccessor{})
)

type Request struct {
	*http.Request
	Params map[string]string
}

func NewRequest(req *http.Request) *Request {
	a := new(Request)
	a.Request = req
	return a
}

func (r *Request) SetParams(p map[string]string) *Request {
	return requestLens.Run(r).Set(p).(*Request)
}

type requestAccessor struct{}

func (s requestAccessor) Get(x AnyVal) AnyVal {
	return x.(*Request).Params
}

func (s requestAccessor) Set(x AnyVal, y AnyVal) AnyVal {
	a := NewRequest(x.(*Request).Request)
	a.Params = y.(map[string]string)
	return a
}
