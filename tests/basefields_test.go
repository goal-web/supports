package tests

import (
	"testing"

	"github.com/goal-web/contracts"
	"github.com/goal-web/supports"
)

type dummyProvider struct{ fields contracts.Fields }
func (d dummyProvider) ToFields() contracts.Fields { return d.fields }

func TestBaseFieldsOptionalGetterPrecedence(t *testing.T) {
    b := &supports.BaseFields{Provider: dummyProvider{fields: contracts.Fields{"a": "x"}}}
    // without OptionalGetter
    if b.StringOptional("a", "def") != "x" { t.Fatalf("field lookup failed") }
    // with OptionalGetter providing override
    b.OptionalGetter = func(key string, defaultValue any) any {
        if key == "a" { return "y" }
        return nil
    }
    if b.StringOptional("a", "def") != "y" { t.Fatalf("optional getter precedence failed") }
}

func TestBaseFieldsOnlyExcept(t *testing.T) {
    b := &supports.BaseFields{Provider: dummyProvider{fields: contracts.Fields{"a": 1, "b": 2}}}
    only := b.Only("a")
    if len(only) != 1 || only["a"] != 1 { t.Fatalf("only failed") }
    except := b.ExceptFields("a")
    if len(except) != 1 || except["b"] != 2 { t.Fatalf("except failed") }
}