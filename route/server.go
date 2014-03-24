package route

import "net/http"

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

func Listen(address string) *http.Server {
	return &http.Server{
		Addr:    address,
		Handler: http.HandlerFunc(handle),
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
}
