package utils

import (
	"errors"
	"github.com/goal-web/contracts"
	"reflect"
	"strings"
)

var (
	UnCallable = errors.New("不可调用的参数！")
)

func GetCallable(arg any) (reflect.Value, error) {
	argValue := reflect.ValueOf(arg)

	if argValue.Kind() == reflect.Func {
		return argValue, nil
	}

	return argValue, UnCallable
}

// IsSameStruct 判断是否同一个结构体
func IsSameStruct(v1, v2 any) bool {
	var (
		f1 reflect.Type
		f2 reflect.Type
		ok bool
	)

	if f1, ok = v1.(reflect.Type); !ok {
		f1 = reflect.TypeOf(v1)
	}

	if f2, ok = v2.(reflect.Type); !ok {
		f2 = reflect.TypeOf(v2)
	}

	return f1.PkgPath() == f2.PkgPath() && f1.Name() == f2.Name()
}

// ToTypes 把变量转换成反射类型
func ToTypes(args ...any) []reflect.Type {
	types := make([]reflect.Type, 0)
	for _, arg := range args {
		types = append(types, reflect.TypeOf(arg))
	}
	return types
}

// IsInstanceIn InstanceIn 判断变量是否是某些类型
func IsInstanceIn(v any, types ...reflect.Type) bool {
	for _, e := range types {
		if IsSameStruct(e, v) {
			return true
		}
	}
	return false
}

// EachStructField 遍历结构体的字段
func EachStructField(value reflect.Value, data any, handler func(reflect.StructField, reflect.Value)) {
	dataType := reflect.TypeOf(data)

	for i := 0; i < dataType.NumField(); i++ {
		handler(dataType.Field(i), value.Field(i))
	}
}

// EachSlice 遍历任意 slice 或者 array
func EachSlice(value reflect.Value, handler func(int, reflect.Value)) {
	sliceLen := value.Len()

	for i := 0; i < sliceLen; i++ {
		handler(i, value.Index(i))
	}
}

// EachMap 遍历任意 map
func EachMap(value reflect.Value, handler func(key reflect.Value, value reflect.Value)) {
	for _, key := range value.MapKeys() {
		handler(key, value.MapIndex(key))
	}
}

// GetTypeKey 获取类型唯一字符串
func GetTypeKey(p reflect.Type) string {
	if p.Kind() == reflect.Ptr {
		p = p.Elem()
	}

	if p.Kind() == reflect.Slice {
		return "[]" + p.Elem().String()
	}

	pkgPath := p.PkgPath()

	return IfString(pkgPath == "", "", pkgPath+".") + p.Name()
}

// NotNil 尽量不要 nil
func NotNil(args ...any) any {
	for _, arg := range args {
		switch argValue := arg.(type) {
		case contracts.InstanceProvider[any]:
			arg = argValue()
		case func() any:
			arg = argValue()
		}
		if arg != nil {
			return arg
		}
	}
	return nil
}

// ParseStructTag 解析结构体的tag
func ParseStructTag(rawTag reflect.StructTag) map[string][]string {
	results := make(map[string][]string, 0)
	for _, tagString := range strings.Split(string(rawTag), " ") {
		tag := strings.Split(tagString, ":")
		if len(tag) > 1 {
			results[tag[0]] = strings.Split(strings.ReplaceAll(tag[1], `"`, ""), ",")
		} else {
			results[tag[0]] = nil
		}
	}
	return results
}

// ToValue 把 interface 转换成指定类型的 reflect.Value
func ToValue(argType reflect.Type, arg any) reflect.Value {
	switch argType.Kind() {
	case reflect.String:
		return reflect.ValueOf(ToString(arg, ""))
	case reflect.Int:
		return reflect.ValueOf(ToInt(arg, 0))
	case reflect.Int64:
		return reflect.ValueOf(ToInt64(arg, 0))
	case reflect.Float64:
		return reflect.ValueOf(ToFloat64(arg, 0))
	case reflect.Float32:
		return reflect.ValueOf(ToFloat(arg, 0))
	case reflect.Bool:
		return reflect.ValueOf(ToBool(arg, false))
	}
	if reflect.TypeOf(arg).ConvertibleTo(argType) {
		return reflect.ValueOf(arg).Convert(argType)
	}
	return reflect.ValueOf(arg)
}
