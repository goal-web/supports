package utils

import "testing"

func TestConversions(t *testing.T) {
    if ToInt("10", 0) != 10 { t.Fatalf("ToInt failed") }
    if ToInt64("20", 0) != 20 { t.Fatalf("ToInt64 failed") }
    if !ToBool("true", false) { t.Fatalf("ToBool failed") }
    if ToString(123, "") != "123" { t.Fatalf("ToString failed") }
}