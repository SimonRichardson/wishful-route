package main

import (
	"fmt"
	. "github.com/SimonRichardson/wishful-route/route"
	. "github.com/SimonRichardson/wishful-route/route/html"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

const port string = "127.0.0.1:8080"

func main() {
	fmt.Println("Booting...")

	server := Serve(Listen(port, Route(
		func() Promise {
			return NotFound("Nope!")
		},
		[]func(x *Request) Option{
			Describe(
				"Get the default route.",
				Get(NewLeft("/"), func(req *Request) Promise {
					return UseTemplate(`Hello World!`).Chain(func(x AnyVal) Monad {
						return Ok(x.(string))
					}).(Promise)
				}),
			),
			Describe(
				"Echo the value sent via parameters.",
				Get(NewRight("/:echo"), func(req *Request) Promise {
					return Ok(fmt.Sprintf("<h1>%s</h1>", req.Params["echo"]))
				}),
			),
		},
	))).UnsafePerform()

	server.(Either).Fold(
		func(x AnyVal) AnyVal {
			fmt.Println("Failed to listen:", x)
			return x
		},
		Identity,
	)
}
