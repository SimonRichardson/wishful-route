package route

import (
	"net/http"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

type Server struct {
	server *http.Server
}

func Listen(address string, route func(x *Request) Promise) IO {
	return NewIO(func() AnyVal {
		return &Server{
			server: &http.Server{
				Addr:    address,
				Handler: http.HandlerFunc(handle(route)),
			},
		}
	})
}

func Serve(s IO) IO {
	return s.Map(func(x AnyVal) AnyVal {
		a := x.(*Server)
		if err := a.server.ListenAndServe(); err != nil {
			return NewLeft(err)
		}
		return NewRight(a)
	}).(IO)
}

func Handle(route func(x *Request) Promise) IO {
	return NewIO(func() AnyVal {
		return http.HandlerFunc(handle(route))
	})
}

func handle(route func(x *Request) Promise) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// I wonder if we should pass a response (monad writer)
		route(NewRequest(r)).Fork(func(x AnyVal) AnyVal {
			result := x.(Result)
			header := w.Header()
			for k, v := range result.Headers {
				header.Set(k, v)
			}
			w.WriteHeader(result.StatusCode)
			w.Write([]byte(result.Body))
			return x
		})
	}
}
