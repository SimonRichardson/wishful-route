package route

import (
	"net/http"
)

type Request struct {
	Request http.Request
}

func NewRequest(request http.Request) *Request {
	return &Request{
		Request: request,
	}
}
