package route

import (
	"net/http"
	. "github.com/SimonRichardson/wishful/wishful"
)

/*
Listen("8080", Route(
    func() Promise {
        return NotFound("Nope!")
    }
), []func(x AnyVal) Option {
    Get('/',  func(req *http.Request) AnyVal {
        return Ok("Hello World!")
    }),
})
*/

func Listen(address string, route func(x AnyVal) AnyVal) *http.Server {
	return &http.Server{
		Addr:    address,
		Handler: http.HandlerFunc(handle(route)),
	}
}

func handle(route func(x AnyVal) AnyVal) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
