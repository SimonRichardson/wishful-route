wishful-route
=============

Wishful http routing.

### Wishful Route

[![Build Status](https://api.travis-ci.org/SimonRichardson/wishful-route.png)](https://travis-ci.org/SimonRichardson/wishful-route)

### Example

This is an example of what a simple route could look like (this is going through
some major workings, so this might be already out of date. See source code for
latest documentation!)

```go
Serve(Listen(address, Route(
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
)))
```
