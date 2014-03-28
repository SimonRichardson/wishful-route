package route

import "net/http"

type Request struct {
	*http.Request
	Params map[string]string
}

func NewRequest(req *http.Request) *Request {
	a := new(Request)
	a.Request = req
	return a
}
