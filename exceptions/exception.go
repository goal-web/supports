package exceptions

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/pkg/errors"
)

func WithError(err error) contracts.Exception {
	if err == nil {
		return nil
	}
	if e, isException := err.(contracts.Exception); isException {
		return e
	}
	return New(err.Error())
}

func WithRecover(err any) contracts.Exception {
	if err == nil {
		return nil
	}
	switch e := err.(type) {
	case contracts.Exception:
		return e
	case string:
		return New(e)
	case fmt.Stringer:
		return New(e.String())
	}
	return New(fmt.Sprintf("%v", err))
}

func WithPrevious(err error, previous contracts.Exception) contracts.Exception {
	if err == nil {
		return nil
	}
	return &Exception{err, previous}
}

func New(err string) contracts.Exception {
	return &Exception{errors.New(err), nil}
}

func Throw(err any) {
	if err != nil {
		panic(WrapException(err))
	}
}

type Exception struct {
	Err      error
	Previous contracts.Exception
}

func (e *Exception) Error() string {
	return e.Err.Error()
}

func (e *Exception) GetPrevious() contracts.Exception {
	return e.Previous
}
