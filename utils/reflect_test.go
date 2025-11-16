package utils

import (
    "reflect"
    "testing"
)

func TestIsSameStruct(t *testing.T) {
    if !IsSameStruct(reflect.TypeOf(0), reflect.TypeOf(int(1))) {
        t.Fatalf("int types should be same struct")
    }
}

type dummy struct{ A int }

func TestToTypesAndIsInstanceIn(t *testing.T) {
    e := dummy{A: 1}
    types := ToTypes(e)
    if len(types) != 1 {
        t.Fatalf("ToTypes length mismatch")
    }
    if !IsInstanceIn(e, types...) {
        t.Fatalf("IsInstanceIn should match")
    }
}

func TestGetTypeKey(t *testing.T) {
    key := GetTypeKey(reflect.TypeOf([]string{}))
    if key == "" {
        t.Fatalf("GetTypeKey should not be empty")
    }
}

func TestToValue(t *testing.T) {
    v := ToValue(reflect.TypeOf(int(0)), "42")
    if !v.IsValid() || v.Int() != 42 {
        t.Fatalf("ToValue int conversion failed: %v", v)
    }
}