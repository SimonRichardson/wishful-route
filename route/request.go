package route

import (
	"net/http"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	lens Lens = Lens{}.AccessorLens(accessor{})
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
	return lens.Run(r).Set(p).(*Request)
}

type accessor struct{}

func (s accessor) Get(x AnyVal) AnyVal {
	return x.(*Request).Params
}

func (s accessor) Set(x AnyVal, y AnyVal) AnyVal {
	a := NewRequest(x.(*Request).Request)
	a.Params = y.(map[string]string)
	return a
}
