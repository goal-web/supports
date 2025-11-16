package tests

import (
	"reflect"
	"testing"

	"github.com/goal-web/supports/exceptions"
	"github.com/goal-web/supports/utils"
)

func TestShouldReport(t *testing.T) {
    // Empty dont-report list: should report
    h := exceptions.DefaultExceptionHandler{DontReportExceptions: []reflect.Type{}}
    if !h.ShouldReport(exceptions.New("x")) { t.Fatalf("should report") }

    // Populate dont-report types correctly: should not report
    e := exceptions.New("y")
    h2 := exceptions.DefaultExceptionHandler{DontReportExceptions: utils.ToTypes(e)}
    if h2.ShouldReport(e) { t.Fatalf("should not report") }
}
