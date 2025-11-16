package exceptions

import (
    "testing"
    "github.com/goal-web/contracts"
)

func TestShouldReport(t *testing.T) {
    h := NewDefaultHandler([]contracts.Exception{})
    if !h.ShouldReport(New("x")) { t.Fatalf("should report") }

    e := New("y")
    h2 := NewDefaultHandler([]contracts.Exception{e})
    if h2.ShouldReport(e) { t.Fatalf("should not report") }
}
