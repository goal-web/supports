package tests

import (
    "testing"

	"github.com/goal-web/supports/commands"
)

func TestParseSignature(t *testing.T) {
    sig := "demo {name} {count?} {--force} {--driver=aes}"
    name, args := commands.ParseSignature(sig)
    if name != "demo" {
        t.Fatalf("name != demo: %s", name)
    }
    if len(args) != 4 {
        t.Fatalf("args len != 4: %d", len(args))
    }
    if args[0].GetName() != "name" || args[0].GetType() != 1 {
        t.Fatalf("required arg mismatch")
    }
    if args[1].GetName() != "count" || args[1].GetType() != 2 {
        t.Fatalf("optional arg mismatch")
    }
    if args[2].GetName() != "force" || args[2].GetType() != 3 {
        t.Fatalf("option mismatch")
    }
    if args[3].GetName() != "driver" || args[3].GetType() != 3 || args[3].GetDefault() != "aes" {
        t.Fatalf("option default mismatch")
    }
}

