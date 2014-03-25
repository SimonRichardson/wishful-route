wishful-route
=============

Wishful http routing.

### Wishful Route

[![Build Status](https://api.travis-ci.org/SimonRichardson/wishful-route.png)](https://travis-ci.org/SimonRichardson/wishful-route)

![](http://cloudfront-assets.reason.com/assets/mc/_external/2013_07/beaker-what-is-this-i-dont-eve.gif)

### Example

This is an example of what a simple route could look like (this is going through
some major workings, so this might be already out of date. See source code for
latest documentation!)

```go
Listen(8080, Route(
    func() Promise {
        return NotFound("Nope!")
    }
), []func(x AnyVal) Option {
    Get('/',  func(req *http.Request) AnyVal {
        return Ok("Hello World!")
    }),
})
```
