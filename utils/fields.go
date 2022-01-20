package utils

import (
	"errors"
	"github.com/goal-web/contracts"
	"reflect"
	"strings"
)

// MergeFields 合并两个 contracts.Fields
func MergeFields(fields contracts.Fields, finalFields contracts.Fields) {
	for key, value := range finalFields {
		fields[key] = value
	}
}

// GetStringField 获取 Fields 中的字符串，会尝试转换类型
func GetStringField(fields contracts.Fields, key string, defaultValues ...string) string {
	if value, existsString := fields[key]; existsString {
		if str, isString := value.(string); isString {
			return str
		}
	}
	return StringOr(defaultValues...)
}

// GetSubField 获取下级 Fields ，如果没有的话，匹配同前缀的放到下级 Fields 中
func GetSubField(fields contracts.Fields, key string, defaultValues ...contracts.Fields) contracts.Fields {

	if subField, isField := fields[key].(contracts.Fields); isField {
		return subField
	}

	if len(defaultValues) > 0 {
		return defaultValues[0]
	}

	subField := make(contracts.Fields)
	prefix := key + "."

	for fieldKey, fieldValue := range fields {
		if strings.HasPrefix(fieldKey, prefix) {
			subField[strings.ReplaceAll(fieldKey, prefix, "")] = fieldValue
		}
	}

	if len(subField) > 0 {
		fields[key] = subField
	}

	return subField
}

// GetInt64Field 获取 Fields 中的 int64，会尝试转换类型
func GetInt64Field(fields contracts.Fields, key string, defaultValues ...int64) int64 {
	var defaultValue int64 = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(int64); isInt {
			return intValue
		}
		return ConvertToInt64(value, defaultValue)
	} else {
		return defaultValue
	}
}

// GetIntField 获取 Fields 中的 int，会尝试转换类型
func GetIntField(fields contracts.Fields, key string, defaultValues ...int) int {
	var defaultValue = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(int); isInt {
			return intValue
		}
		return int(ConvertToInt64(value, int64(defaultValue)))
	} else {
		return defaultValue
	}
}

// GetFloatField 获取 Fields 中的 float32，会尝试转换类型
func GetFloatField(fields contracts.Fields, key string, defaultValues ...float32) float32 {
	var defaultValue float32 = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(float32); isInt {
			return intValue
		}
		return ConvertToFloat(value, defaultValue)
	} else {
		return defaultValue
	}
}

// GetFloat64Field 获取 Fields 中的 float64，会尝试转换类型
func GetFloat64Field(fields contracts.Fields, key string, defaultValues ...float64) float64 {
	var defaultValue float64 = 0
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if value, existsValue := fields[key]; existsValue {
		if intValue, isInt := value.(float64); isInt {
			return intValue
		}
		return ConvertToFloat64(value, defaultValue)
	} else {
		return defaultValue
	}
}

// GetBoolField 获取 Fields 中的 bool，会尝试转换类型
func GetBoolField(fields contracts.Fields, key string, defaultValues ...bool) bool {
	var defaultValue = false
	if len(defaultValues) > 0 {
		defaultValue = defaultValues[0]
	}
	if fieldValue, existsValue := fields[key]; existsValue {
		return ConvertToBool(fieldValue, defaultValue)
	}
	return defaultValue
}

// ConvertToFields 尝试把一个变量转换成 Fields 类型
func ConvertToFields(anyValue interface{}) (contracts.Fields, error) {
	fields := contracts.Fields{}
	switch paramValue := anyValue.(type) {
	case contracts.Fields:
		fields = paramValue
	case map[string]interface{}:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]int:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]float64:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]string:
		for key, value := range paramValue {
			fields[key] = value
		}
	case map[string]bool:
		for key, value := range paramValue {
			fields[key] = value
		}
	default:
		paramType := reflect.ValueOf(anyValue)

		switch paramType.Kind() {
		case reflect.Struct: // 结构体
			EachStructField(anyValue, func(field reflect.StructField, value reflect.Value) {
				fields[SnakeString(field.Name)] = value.Interface()
			})
		case reflect.Map: // 自定义的 map
			if paramType.Type().Key().Kind() == reflect.String {
				for _, key := range paramType.MapKeys() {
					name := key.String()
					fields[name] = paramType.MapIndex(key).Interface()
				}
			} else {
				return nil, errors.New("不支持 string 以外的类型作为 key 的 map")
			}
		default:
			return nil, errors.New("不支持转 contracts.Fields 的类型： " + paramType.String())
		}
	}
	return fields, nil
}
