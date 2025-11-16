package exceptions

import (
    "testing"
)

func TestWithErrorAndRecover(t *testing.T) {
    if WithError(nil) != nil { t.Fatalf("nil error unexpected") }
    e := WithError(New("x"))
    if e == nil { t.Fatalf("expected exception") }

    if WithRecover(nil) != nil { t.Fatalf("nil recover unexpected") }
    if WithRecover("err").Error() == "" { t.Fatalf("string recover empty") }
}

func TestErrorConcat(t *testing.T) {
    e := &Exception{Err: New("a").(*Exception).Err, Previous: New("b")}
    if e.Error() != "a.b" { t.Fatalf("concat mismatch: %s", e.Error()) }
}
