package route

import (
	"encoding/json"
	. "github.com/SimonRichardson/wishful/useful"
	. "github.com/SimonRichardson/wishful/wishful"
	"net/url"
)

var (
	EitherPromise EitherT = NewEitherT(Promise{})
)

func JsonParse(m interface{}) EitherT {
	promise := Promise.Of(handleException(func(data AnyVal) Either {
		if err := json.Unmarshal(data, &m); err != nil {
			return NewLeft(err)
		}
		return NewRight(m)
	}))
	return EitherPromise.From(promise)
}

func QueryParse(u string) EitherT {
	promise := Promise.Of(handleException(func(data AnyVal) Either {
		u, err := url.Parse(u)
		if err != nil {
			return NewLeft(err)
		}
		return NewRight(u.Query())
	}))
	return EitherPromise.From(promise)
}
