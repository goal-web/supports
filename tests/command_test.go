package tests

import (
	"testing"

	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/commands"
)

type fakeArgs struct {
	options map[string]any
	args    []string
}

func (f *fakeArgs) Get(key string) any { return f.options[key] }
func (f *fakeArgs) Optional(key string, defaultValue any) any {
	if v, ok := f.options[key]; ok {
		return v
	}
	return defaultValue
}
func (f *fakeArgs) GetArg(index int) string {
	if index < len(f.args) {
		return f.args[index]
	}
	return ""
}
func (f *fakeArgs) GetArgs() []string               { return f.args }
func (f *fakeArgs) SetOption(key string, value any) { f.options[key] = value }
func (f *fakeArgs) Exists(key string) bool          { _, ok := f.options[key]; return ok }
func (f *fakeArgs) StringArrayOption(key string, defaultValue []string) []string {
	if v, ok := f.options[key].([]string); ok {
		return v
	}
	return defaultValue
}
func (f *fakeArgs) IntArrayOption(key string, defaultValue []int) []int {
	if v, ok := f.options[key].([]int); ok {
		return v
	}
	return defaultValue
}
func (f *fakeArgs) Int64ArrayOption(key string, defaultValue []int64) []int64 {
	if v, ok := f.options[key].([]int64); ok {
		return v
	}
	return defaultValue
}
func (f *fakeArgs) FloatArrayOption(key string, defaultValue []float32) []float32 {
	if v, ok := f.options[key].([]float32); ok {
		return v
	}
	return defaultValue
}
func (f *fakeArgs) Float64ArrayOption(key string, defaultValue []float64) []float64 {
    if v, ok := f.options[key].([]float64); ok {
        return v
    }
    return defaultValue
}

func (f *fakeArgs) GetString(key string) string {
    if v, ok := f.Get(key).(string); ok {
        return v
    }
    return ""
}

func (f *fakeArgs) GetInt64(key string) int64 {
    if v, ok := f.Get(key).(int64); ok {
        return v
    }
    if v, ok := f.Get(key).(int); ok {
        return int64(v)
    }
    return 0
}

func (f *fakeArgs) GetInt32(key string) int32 {
    if v, ok := f.Get(key).(int32); ok {
        return v
    }
    if v, ok := f.Get(key).(int); ok {
        return int32(v)
    }
    return 0
}

func (f *fakeArgs) GetInt16(key string) int16 {
    if v, ok := f.Get(key).(int16); ok {
        return v
    }
    if v, ok := f.Get(key).(int); ok {
        return int16(v)
    }
    return 0
}

func (f *fakeArgs) GetInt8(key string) int8 {
    if v, ok := f.Get(key).(int8); ok {
        return v
    }
    if v, ok := f.Get(key).(int); ok {
        return int8(v)
    }
    return 0
}

func (f *fakeArgs) GetInt(key string) int {
    if v, ok := f.Get(key).(int); ok {
        return v
    }
    return 0
}

func (f *fakeArgs) GetUInt64(key string) uint64 {
    if v, ok := f.Get(key).(uint64); ok {
        return v
    }
    return 0
}

func (f *fakeArgs) GetUInt32(key string) uint32 {
    if v, ok := f.Get(key).(uint32); ok {
        return v
    }
    return 0
}

func (f *fakeArgs) GetUInt16(key string) uint16 {
    if v, ok := f.Get(key).(uint16); ok {
        return v
    }
    return 0
}

func (f *fakeArgs) GetUInt8(key string) uint8 {
    if v, ok := f.Get(key).(uint8); ok {
        return v
    }
    return 0
}

func (f *fakeArgs) GetUInt(key string) uint {
    if v, ok := f.Get(key).(uint); ok {
        return v
    }
    return 0
}

func (f *fakeArgs) GetFloat64(key string) float64 {
    if v, ok := f.Get(key).(float64); ok {
        return v
    }
    return 0
}

func (f *fakeArgs) GetFloat(key string) float32 {
    if v, ok := f.Get(key).(float32); ok {
        return v
    }
    return 0
}

func (f *fakeArgs) GetBool(key string) bool {
    if v, ok := f.Get(key).(bool); ok {
        return v
    }
    return false
}

func (f *fakeArgs) StringOptional(key string, defaultValue string) string {
    v := f.GetString(key)
    if v == "" {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) Int64Optional(key string, defaultValue int64) int64 {
    v := f.GetInt64(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) Int32Optional(key string, defaultValue int32) int32 {
    v := f.GetInt32(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) Int16Optional(key string, defaultValue int16) int16 {
    v := f.GetInt16(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) Int8Optional(key string, defaultValue int8) int8 {
    v := f.GetInt8(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) IntOptional(key string, defaultValue int) int {
    v := f.GetInt(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) UInt64Optional(key string, defaultValue uint64) uint64 {
    v := f.GetUInt64(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) UInt32Optional(key string, defaultValue uint32) uint32 {
    v := f.GetUInt32(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) UInt16Optional(key string, defaultValue uint16) uint16 {
    v := f.GetUInt16(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) UInt8Optional(key string, defaultValue uint8) uint8 {
    v := f.GetUInt8(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) UIntOptional(key string, defaultValue uint) uint {
    v := f.GetUInt(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) Float64Optional(key string, defaultValue float64) float64 {
    v := f.GetFloat64(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) FloatOptional(key string, defaultValue float32) float32 {
    v := f.GetFloat(key)
    if v == 0 {
        return defaultValue
    }
    return v
}

func (f *fakeArgs) BoolOptional(key string, defaultValue bool) bool {
    v := f.GetBool(key)
    if !v {
        return defaultValue
    }
    return v
}

func TestInjectArguments(t *testing.T) {
	cmd := commands.Base("demo {name} {count?} {--force} {--driver=aes}", "")
	args := &fakeArgs{options: map[string]any{}, args: []string{"alice"}}
	if err := cmd.InjectArguments(cmd.GetArgs(), args); err != nil {
		t.Fatalf("inject error: %v", err)
	}
	if args.Get("name") != "alice" {
		t.Fatalf("name not injected")
	}
	if args.Get("count") != nil {
		t.Fatalf("count default unexpected")
	}
	if args.Get("driver") != "aes" {
		t.Fatalf("driver default missing")
	}
}

func TestInjectArgumentsMissingRequired(t *testing.T) {
	cmd := commands.Base("demo {name} {count?}", "")
	args := &fakeArgs{options: map[string]any{}, args: []string{""}}
	if err := cmd.InjectArguments(cmd.GetArgs(), args); err == nil {
		t.Fatalf("expected error")
	}
}

var _ contracts.CommandArguments = (*fakeArgs)(nil)
