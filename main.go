package main

import (
	"fmt"
	. "github.com/SimonRichardson/wishful-route/route"
	. "github.com/SimonRichardson/wishful-route/route/jsonp"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

const port string = "127.0.0.1:8080"

type Message struct {
	Message string
}

type Echo struct {
	Echo string
}

func NewMessage(a string) Message {
	return Message{
		Message: a,
	}
}

func NewEcho(a string) Echo {
	return Echo{
		Echo: a,
	}
}

func main() {
	fmt.Println("Booting...")

	server := Serve(Listen(port, Route(
		func() Promise {
			return NotFound("callback", NewMessage("Nope!"))
		},
		[]func(x *Request) Option{
			Describe(
				"Get the default route.",
				Get(NewLeft("/"), func(req *Request) Promise {
					return Ok("callback", NewMessage("Hello World!"))
				}),
			),
			Describe(
				"Echo the value sent via parameters.",
				Get(NewRight("/:echo"), func(req *Request) Promise {
					return Ok("callback", NewEcho(req.Params["echo"]))
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
