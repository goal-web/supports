package utils

import (
	"fmt"
	"strconv"
)

// ToInt64 把能转换成 int64 的值转换成 int64
func ToInt64(rawValue any, defaultValue int64) int64 {
	switch value := rawValue.(type) {
	case int64:
		return value
	case int:
		return int64(value)
	case uint:
		return int64(value)
	case uint32:
		return int64(value)
	case uint8:
		return int64(value)
	case uint16:
		return int64(value)
	case uint64:
		return int64(value)
	case int8:
		return int64(value)
	case int16:
		return int64(value)
	case int32:
		return int64(value)
	case float64:
		return int64(value)
	case float32:
		return int64(value)
	case []rune:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return i64
	case []byte:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return i64
	case string:
		i64, _ := strconv.ParseInt(value, 10, 64)
		return i64
	}

	return defaultValue
}

// uint
func ToUInt64(rawValue any, defaultValue uint64) uint64 {
	switch value := rawValue.(type) {
	case int64:
		return uint64(value)
	case int:
		return uint64(value)
	case uint:
		return uint64(value)
	case uint32:
		return uint64(value)
	case uint8:
		return uint64(value)
	case uint16:
		return uint64(value)
	case uint64:
		return uint64(value)
	case int8:
		return uint64(value)
	case int16:
		return uint64(value)
	case int32:
		return uint64(value)
	case float64:
		return uint64(value)
	case float32:
		return uint64(value)
	case []rune:
		i64, _ := strconv.ParseUint(string(value), 10, 64)
		return i64
	case []byte:
		i64, _ := strconv.ParseUint(string(value), 10, 64)
		return i64
	case string:
		i64, _ := strconv.ParseUint(value, 10, 64)
		return i64
	}

	return defaultValue
}

// ToInt 把能转换成 int 的值转换成 int
func ToInt(rawValue any, defaultValue int) int {
	switch value := rawValue.(type) {
	case int64:
		return int(value)
	case int:
		return value
	case uint:
		return int(value)
	case uint32:
		return int(value)
	case uint8:
		return int(value)
	case uint16:
		return int(value)
	case uint64:
		return int(value)
	case int8:
		return int(value)
	case int16:
		return int(value)
	case int32:
		return int(value)
	case float64:
		return int(value)
	case float32:
		return int(value)
	case []rune:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return int(i64)
	case []byte:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return int(i64)
	case string:
		i64, _ := strconv.ParseInt(value, 10, 32)
		return int(i64)
	}

	return defaultValue
}

// uint
func ToUInt(rawValue any, defaultValue uint) uint {
	switch value := rawValue.(type) {
	case int64:
		return uint(value)
	case int:
		return uint(value)
	case uint:
		return uint(value)
	case uint32:
		return uint(value)
	case uint8:
		return uint(value)
	case uint16:
		return uint(value)
	case uint64:
		return uint(value)
	case int8:
		return uint(value)
	case int16:
		return uint(value)
	case int32:
		return uint(value)
	case float64:
		return uint(value)
	case float32:
		return uint(value)
	case []rune:
		i64, _ := strconv.ParseUint(string(value), 10, 32)
		return uint(i64)
	case []byte:
		i64, _ := strconv.ParseUint(string(value), 10, 32)
		return uint(i64)
	case string:
		i64, _ := strconv.ParseUint(value, 10, 32)
		return uint(i64)
	}

	return defaultValue
}

// ToInt32 把能转换成 int 的值转换成 int
func ToInt32(rawValue any, defaultValue int32) int32 {
	switch value := rawValue.(type) {
	case int64:
		return int32(value)
	case int:
		return int32(value)
	case uint:
		return int32(value)
	case uint32:
		return int32(value)
	case uint8:
		return int32(value)
	case uint16:
		return int32(value)
	case uint64:
		return int32(value)
	case int8:
		return int32(value)
	case int16:
		return int32(value)
	case int32:
		return value
	case float64:
		return int32(value)
	case float32:
		return int32(value)
	case []rune:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return int32(i64)
	case []byte:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return int32(i64)
	case string:
		i64, _ := strconv.ParseInt(value, 10, 32)
		return int32(i64)
	}

	return defaultValue
}

// uint
func ToUInt32(rawValue any, defaultValue uint32) uint32 {
	switch value := rawValue.(type) {
	case int64:
		return uint32(value)
	case int:
		return uint32(value)
	case uint:
		return uint32(value)
	case uint32:
		return uint32(value)
	case uint8:
		return uint32(value)
	case uint16:
		return uint32(value)
	case uint64:
		return uint32(value)
	case int8:
		return uint32(value)
	case int16:
		return uint32(value)
	case int32:
		return uint32(value)
	case float64:
		return uint32(value)
	case float32:
		return uint32(value)
	case []rune:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return uint32(i64)
	case []byte:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return uint32(i64)
	case string:
		i64, _ := strconv.ParseInt(value, 10, 32)
		return uint32(i64)
	}

	return defaultValue
}

// ToInt8 把能转换成 int 的值转换成 int
func ToInt8(rawValue any, defaultValue int8) int8 {
	switch value := rawValue.(type) {
	case int64:
		return int8(value)
	case int:
		return int8(value)
	case uint:
		return int8(value)
	case uint32:
		return int8(value)
	case uint8:
		return int8(value)
	case uint16:
		return int8(value)
	case uint64:
		return int8(value)
	case int8:
		return int8(value)
	case int16:
		return int8(value)
	case int32:
		return int8(value)
	case float64:
		return int8(value)
	case float32:
		return int8(value)
	case []rune:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return int8(i64)
	case []byte:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return int8(i64)
	case string:
		i64, _ := strconv.ParseInt(value, 10, 32)
		return int8(i64)
	}

	return defaultValue
}

// uint
func ToUInt8(rawValue any, defaultValue uint8) uint8 {
	switch value := rawValue.(type) {
	case int64:
		return uint8(value)
	case int:
		return uint8(value)
	case uint:
		return uint8(value)
	case uint32:
		return uint8(value)
	case uint8:
		return uint8(value)
	case uint16:
		return uint8(value)
	case uint64:
		return uint8(value)
	case int8:
		return uint8(value)
	case int16:
		return uint8(value)
	case int32:
		return uint8(value)
	case float64:
		return uint8(value)
	case float32:
		return uint8(value)
	case []rune:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return uint8(i64)
	case []byte:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return uint8(i64)
	case string:
		i64, _ := strconv.ParseInt(value, 10, 32)
		return uint8(i64)
	}

	return defaultValue
}

// ToInt16 把能转换成 int 的值转换成 int
func ToInt16(rawValue any, defaultValue int16) int16 {
	switch value := rawValue.(type) {
	case int64:
		return int16(value)
	case int:
		return int16(value)
	case uint:
		return int16(value)
	case uint32:
		return int16(value)
	case uint8:
		return int16(value)
	case uint16:
		return int16(value)
	case uint64:
		return int16(value)
	case int8:
		return int16(value)
	case int16:
		return int16(value)
	case int32:
		return int16(value)
	case float64:
		return int16(value)
	case float32:
		return int16(value)
	case []rune:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return int16(i64)
	case []byte:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return int16(i64)
	case string:
		i64, _ := strconv.ParseInt(value, 10, 32)
		return int16(i64)
	}

	return defaultValue
}

// uint
func ToUInt16(rawValue any, defaultValue uint16) uint16 {
	switch value := rawValue.(type) {
	case int64:
		return uint16(value)
	case int:
		return uint16(value)
	case uint:
		return uint16(value)
	case uint32:
		return uint16(value)
	case uint8:
		return uint16(value)
	case uint16:
		return uint16(value)
	case uint64:
		return uint16(value)
	case int8:
		return uint16(value)
	case int16:
		return uint16(value)
	case int32:
		return uint16(value)
	case float64:
		return uint16(value)
	case float32:
		return uint16(value)
	case []rune:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return uint16(i64)
	case []byte:
		i64, _ := strconv.ParseInt(string(value), 10, 32)
		return uint16(i64)
	case string:
		i64, _ := strconv.ParseInt(value, 10, 32)
		return uint16(i64)
	}

	return defaultValue
}

// ToFloat64 把能转换成 float64 的值转换成 float64
func ToFloat64(rawValue any, defaultValue float64) float64 {
	switch value := rawValue.(type) {
	case float64:
		return value
	case int64:
		return float64(value)
	case uint:
		return float64(value)
	case uint32:
		return float64(value)
	case uint8:
		return float64(value)
	case uint16:
		return float64(value)
	case uint64:
		return float64(value)
	case int:
		return float64(value)
	case int8:
		return float64(value)
	case int16:
		return float64(value)
	case int32:
		return float64(value)
	case float32:
		return float64(value)
	case bool:
		if value {
			return 1
		} else {
			return 0
		}
	case []rune:
		f64, _ := strconv.ParseFloat(string(value), 64)
		return f64
	case []byte:
		f64, _ := strconv.ParseFloat(string(value), 64)
		return f64
	case string:
		f64, _ := strconv.ParseFloat(value, 64)
		return f64
	}

	return defaultValue
}

// ToFloat 把能转换成 float32 的值转换成 float32
func ToFloat(rawValue any, defaultValue float32) float32 {
	switch value := rawValue.(type) {
	case float64:
		return float32(value)
	case uint:
		return float32(value)
	case uint32:
		return float32(value)
	case uint8:
		return float32(value)
	case uint16:
		return float32(value)
	case uint64:
		return float32(value)
	case int64:
		return float32(value)
	case int:
		return float32(value)
	case int8:
		return float32(value)
	case int16:
		return float32(value)
	case int32:
		return float32(value)
	case float32:
		return value
	case bool:
		if value {
			return 1
		} else {
			return 0
		}
	case string:
		f64, _ := strconv.ParseFloat(value, 32)
		return float32(f64)
	case []rune:
		f64, _ := strconv.ParseFloat(string(value), 32)
		return float32(f64)
	case []byte:
		f64, _ := strconv.ParseFloat(string(value), 32)
		return float32(f64)
	}

	return defaultValue
}

// ToBool 把能转换成 bool 的值转换成 bool
func ToBool(rawValue any, defaultValue bool) bool {
	switch value := rawValue.(type) {
	case bool:
		return value
	case string:
		switch value {
		case "false", "(false)", "0", "":
			return false
		case "true", "(true)", "1":
			return true
		}
	case float64, int, int64, int8, float32:
		return ToInt64(value, 0) > 0 || defaultValue
	}

	return defaultValue
}

// ToString 把能转换成 string 的值转换成 string
func ToString(rawValue any, defaultValue string) string {
	switch value := rawValue.(type) {
	case bool:
		return IfString(value, "true", "false")
	case string:
		return value
	case []byte:
		return string(value)
	case []rune:
		return string(value)
	case fmt.Stringer:
		return value.String()
	case int, int64, int8, int32, int16, uint16, uint, uint8, uint32, uint64:
		return fmt.Sprintf("%d", value)
	case float32, float64:
		return fmt.Sprintf("%f", value)
		//case any:
		//	return fmt.Sprintf("%v", value)
	}

	return defaultValue
}
