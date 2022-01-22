package exceptions

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/pkg/errors"
)

func WithError(err error, fields contracts.Fields) contracts.Exception {
	if e, isException := err.(contracts.Exception); isException {
		return e
	}
	return New(err.Error(), fields)
}

func WithRecover(err interface{}, fields contracts.Fields) contracts.Exception {
	switch e := err.(type) {
	case contracts.Exception:
		return e
	case error:
		return WithError(e, fields)
	case string:
		return New(e, fields)
	case fmt.Stringer:
		return New(e.String(), fields)
	}
	return New(fmt.Sprintf("%v", err), fields)
}

func WithPrevious(err error, fields contracts.Fields, previous contracts.Exception) Exception {
	return Exception{err, fields, previous}
}

func New(err string, fields contracts.Fields) Exception {
	return Exception{errors.New(err), fields, nil}
}

func Throw(err interface{}) {
	if err != nil {
		panic(ResolveException(err))
	}
}

type Exception struct {
	error
	fields   contracts.Fields
	previous contracts.Exception
}

func (e Exception) Fields() contracts.Fields {
	return e.fields
}
