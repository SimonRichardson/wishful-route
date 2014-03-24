package route

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
)

var (
	EitherPromise EitherT = NewEitherT(Promise{})
)

func from(x Either) EitherT {
	return EitherPromise.From(Promise{}.Of(x))
}

func parseJson(raw []byte, val AnyVal) Either {
	if err := json.Unmarshal(raw, &val); err != nil {
		return NewLeft(err)
	}
	return NewRight(val)
}

func parseQuery(raw string) Either {
	u, err := url.ParseQuery(raw)
	if err != nil {
		return NewLeft(err)
	}
	return NewRight(u)
}

func JsonParse(val AnyVal) func(raw []byte) EitherT {
	return func(raw []byte) EitherT {
		return from(handleException(func(x AnyVal) Either {
			return parseJson(x.([]byte), val)
		})(raw))
	}
}

func QueryParse(raw string) EitherT {
	return from(handleException(func(x AnyVal) Either {
		return parseQuery(x.(string))
	})(raw))
}

func ReadBody(req *http.Request) EitherT {
	c := req.Header.Get("Content-Length")
	length, err := strconv.Atoi(c)
	if err != nil {
		return from(NewLeft(err))
	}

	reader := io.LimitReader(req.Body, int64(length))
	b, e := ioutil.ReadAll(reader)

	if e != nil {
		return from(NewLeft(e))
	}

	if len(b) != length {
		err := errors.New("http: Body length mismatch")
		return from(NewLeft(err))
	}

	return from(NewRight(string(b)))
}

func Json(val AnyVal, req *http.Request) EitherT {
	return ReadBody(req).Chain(func(x AnyVal) Monad {
		return JsonParse(val)(x.([]byte))
	}).(EitherT)
}

func Query(req *http.Request) EitherT {
	return ReadBody(req).Chain(func(x AnyVal) Monad {
		return QueryParse(x.(string))
	}).(EitherT)
}

func Raw(req *http.Request) EitherT {
	return ReadBody(req)
}
