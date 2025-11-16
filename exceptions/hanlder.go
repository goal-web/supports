package exceptions

import (
	"reflect"

	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/logs"
	"github.com/goal-web/supports/utils"
)

type DefaultExceptionHandler struct {
	DontReportExceptions []reflect.Type
}

func NewDefaultHandler(dontReportExceptions []contracts.Exception) DefaultExceptionHandler {
	return DefaultExceptionHandler{utils.ToTypes(dontReportExceptions)}
}

func (handler DefaultExceptionHandler) Handle(exception contracts.Exception) (result any) {
	logs.WithException(exception).
		WithField("exception", reflect.TypeOf(exception).String()).
		Error("DefaultExceptionHandler")

	return
}

func (handler DefaultExceptionHandler) Report(exception contracts.Exception) {
}

func (handler DefaultExceptionHandler) ShouldReport(exception contracts.Exception) bool {
	return !utils.IsInstanceIn(exception, handler.DontReportExceptions...)
}
