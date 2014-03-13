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

func parseJson(raw []byte, val AnyVal) Either {
	if err := json.Unmarshal(raw, &val); err != nil {
		return NewLeft(err)
	}
	return NewRight(val)
}

func parseQuery(raw string) Either {
	u, err := url.Parse(raw)
	if err != nil {
		return NewLeft(err)
	}
	return NewRight(u.Query())
}

func JsonParse(raw []byte, val AnyVal) EitherT {
	promise := Promise{}.Of(handleException(func(x AnyVal) Either {
		return parseJson(x.([]byte), val)
	})(raw))
	return EitherPromise.From(promise)
}

func QueryParse(raw string) EitherT {
	promise := Promise{}.Of(handleException(func(x AnyVal) Either {
		return parseQuery(x.(string))
	})(raw))
	return EitherPromise.From(promise)
}
