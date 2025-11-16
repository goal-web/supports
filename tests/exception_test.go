package tests

import (
    "testing"

	"github.com/goal-web/supports/exceptions"
)

func TestWithErrorAndRecover(t *testing.T) {
    if exceptions.WithError(nil) != nil { t.Fatalf("nil error unexpected") }
    e := exceptions.WithError(exceptions.New("x"))
    if e == nil { t.Fatalf("expected exception") }

    if exceptions.WithRecover(nil) != nil { t.Fatalf("nil recover unexpected") }
    if exceptions.WithRecover("err").Error() == "" { t.Fatalf("string recover empty") }
}

func TestErrorConcat(t *testing.T) {
    e := &exceptions.Exception{Err: exceptions.New("a").(*exceptions.Exception).Err, Previous: exceptions.New("b")}
    if e.Error() != "a.b" { t.Fatalf("concat mismatch: %s", e.Error()) }
}
