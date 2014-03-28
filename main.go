package main

import (
	"fmt"
	. "github.com/SimonRichardson/wishful-route/route"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

const port string = "127.0.0.1:8080"

func main() {
	fmt.Println("Booting...")

	Serve(Listen(port, Route(
		func() Promise {
			return NotFound("Nope!")
		},
		[]func(x *Request) Option{
			Get("/", func(req *Request) Promise {
				return Ok("Hello World!")
			}),
			Get("/:echo", func(req *Request) Promise {
				return Ok(fmt.Sprintf("%s", req.Params["echo"]))
			}),
		},
	))).(Either).Fold(
		func(x AnyVal) AnyVal {
			fmt.Println("Failed to listen:", x)
			return x
		},
		Identity,
	)
}
