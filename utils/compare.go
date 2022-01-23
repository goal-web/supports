package utils

import (
	"fmt"
	"reflect"
)

type CompareHandler func(comparable interface{}, arg interface{}) bool

func Compare(comparable interface{}, operator string, arg interface{}) bool {
	switch operator {
	case "=", "eq":
		return IsEqual(comparable, arg)
	case ">=", "gte":
		return IsGte(comparable, arg)
	case ">", "gt":
		return IsGt(comparable, arg)
	case "<", "lt":
		return IsLt(comparable, arg)
	case "<=", "lte":
		return IsLte(comparable, arg)
	case "in":
		return IsIn(comparable, arg)
	case "not in":
		return IsNotIn(comparable, arg)
	}
	return false
}

// IsEqual 等于 =
func IsEqual(comparable interface{}, arg interface{}) bool {
	comparableType := reflect.TypeOf(comparable)
	argType := reflect.TypeOf(arg)

	if comparableType.Comparable() && argType.Comparable() && comparable == arg {
		return true
	}

	switch comparableType.Kind() {
	case reflect.Bool:
		return comparable.(bool) == ConvertToBool(arg, !comparable.(bool)) // 若不能转换成bool，则默认不相等
	case reflect.String:
		return comparable.(string) == ConvertToString(arg, fmt.Sprintf("%v", arg))
	case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint64, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Float64, reflect.Float32:
		return ConvertToFloat64(comparable, 0) == ConvertToFloat64(arg, 0)
	// 判断结构体是否相等
	case reflect.Struct:
		// 不是同一种结构体
		if !IsSameStruct(comparableType, arg) {
			return false
		}

		argValue := reflect.ValueOf(arg)
		isSame := true
		EachStructField(reflect.ValueOf(comparable), comparable, func(field reflect.StructField, value reflect.Value) {
			if !field.IsExported() {
				isSame = false
			}

			if isSame && IsEqual(value.Interface(), argValue.FieldByName(field.Name).Interface()) == false {
				isSame = false
			}
		})
		return isSame
	case reflect.Array, reflect.Slice:
		comparableValue := reflect.ValueOf(comparable)
		argValue := reflect.ValueOf(arg)
		if comparableValue.Len() != argValue.Len() {
			return false
		}

		isSame := true
		EachSlice(comparableValue, func(index int, value reflect.Value) {
			if isSame && IsEqual(value.Interface(), argValue.Index(index).Interface()) == false {
				isSame = false
			}
		})
		return isSame
	}
	return false
}

// IsIn 存在 in
func IsIn(comparable interface{}, arg interface{}) (result bool) {
	argValue := reflect.ValueOf(arg)
	if !IsArray(argValue) {
		return false
	}

	EachSlice(argValue, func(index int, value reflect.Value) {
		if result == false && IsEqual(comparable, value.Interface()) {
			result = true
		}
	})
	return
}

// IsNotIn 不存在 not in
func IsNotIn(comparable interface{}, arg interface{}) (result bool) {
	argValue := reflect.ValueOf(arg)
	if !IsArray(argValue) {
		return false
	}
	EachSlice(argValue, func(index int, value reflect.Value) {
		if result == false && IsEqual(comparable, value.Interface()) {
			result = true
		}
	})
	return !result
}

// IsLt 小于 <
func IsLt(comparable interface{}, arg interface{}) bool {
	return ConvertToFloat64(comparable, 0) < ConvertToFloat64(arg, 0)
}

// IsLte 小于等于 <=
func IsLte(comparable interface{}, arg interface{}) bool {
	return ConvertToFloat64(comparable, 0) <= ConvertToFloat64(arg, 0)
}

// IsGt 大于 >
func IsGt(comparable interface{}, arg interface{}) bool {
	return ConvertToFloat64(comparable, 0) > ConvertToFloat64(arg, 0)
}

// IsGte 大于等于
func IsGte(comparable interface{}, arg interface{}) bool {
	return ConvertToFloat64(comparable, 0) >= ConvertToFloat64(arg, 0)
}

// IsArray 是否是数组或者 slice
func IsArray(comparable interface{}) bool {
	switch value := comparable.(type) {
	case reflect.Type:
		return value.Kind() == reflect.Slice || value.Kind() == reflect.Array
	case reflect.Value:
		return value.Kind() == reflect.Slice || value.Kind() == reflect.Array
	}
	comparableType := reflect.TypeOf(comparable)
	return comparableType.Kind() == reflect.Slice || comparableType.Kind() == reflect.Array
}
