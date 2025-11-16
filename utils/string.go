package utils

import (
	"fmt"
	"strings"
)

// SubString 切割字符串
func SubString(str string, start, num int) string {
	runes := []rune(str)
	strLen := len(runes)
	if start >= strLen {
		return ""
	}
	if num < 0 {
		return string(runes[start : strLen+num])
	}
	if start+num >= strLen || num == 0 {
		return string(runes[start:])
	}
	return string(runes[start : start+num])
}

// IfString 类似三目运算
func IfString(condition bool, str1 string, otherStr ...string) string {
	if condition {
		return str1
	}
	return StringOr(otherStr...)
}

// Ifi 类似三目运算，返回第一个不是 nil 的值
func Ifi(condition bool, value any, others ...any) any {
	if condition {
		return value
	}
	return NotNil(others...)
}

// StringOr 尽量不返回空字符串
func StringOr(otherStr ...string) string {
	for _, str := range otherStr {
		if str != "" {
			return str
		}
	}
	return ""
}

// IntOr 尽量不返回0值
func IntOr(otherInt ...int) int {
	for _, i := range otherInt {
		if i != 0 {
			return i
		}
	}
	return 0
}

// SnakeString 蛇形字符串
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写
	return strings.ToLower(string(data[:]))
}

// CamelString 驼峰
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if !k && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || !k) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

// JoinStringerArray 连接fmt.Stringer数组，类似 strings.Join
func JoinStringerArray(arr []fmt.Stringer, sep string) (result string) {
	for index, stringer := range arr {
		if index == 0 {
			result = stringer.String()
		} else {
			result = fmt.Sprintf("%s%s%s", result, sep, stringer.String())
		}
	}
	return
}

// JoinIntArray 连接 int 数组，类似 strings.Join
func JoinIntArray(arr []int, sep string) (result string) {
	for index, num := range arr {
		if index == 0 {
			result = fmt.Sprintf("%d", num)
		} else {
			result = fmt.Sprintf("%s%s%d", result, sep, num)
		}
	}

	return
}

// JoinInt64Array 连接 int64 数组，类似 strings.Join
func JoinInt64Array(arr []int64, sep string) (result string) {
	for index, num := range arr {
		if index == 0 {
			result = fmt.Sprintf("%d", num)
		} else {
			result = fmt.Sprintf("%s%s%d", result, sep, num)
		}
	}
	return
}

// JoinFloatArray 连接 float32 数组，类似 strings.Join
func JoinFloatArray(arr []float32, sep string) (result string) {
	for index, num := range arr {
		if index == 0 {
			result = fmt.Sprintf("%f", num)
		} else {
			result = fmt.Sprintf("%s%s%f", result, sep, num)
		}
	}
	return
}

// JoinFloat64Array 连接 float64 数组，类似 strings.Join
func JoinFloat64Array(arr []float64, sep string) (result string) {
	for index, num := range arr {
		if index == 0 {
			result = fmt.Sprintf("%f", num)
		} else {
			result = fmt.Sprintf("%s%s%f", result, sep, num)
		}
	}
	return
}

// JoinInterfaceArray 连接 interface 数组，类似 strings.Join
func JoinInterfaceArray(arr []any, sep string) (result string) {
	for index, v := range arr {
		if index == 0 {
			result = fmt.Sprintf("%v", v)
		} else {
			result = fmt.Sprintf("%s%s%v", result, sep, v)
		}
	}
	return
}

// MakeSymbolArray 创建一个有指定字符组成的数组
func MakeSymbolArray(symbol string, num int) (result []string) {
	for i := 0; i < num; i++ {
		result = append(result, symbol)
	}
	return
}

// StringArray2InterfaceArray 把字符串数组转成 any 数组
func StringArray2InterfaceArray(args []string) (result []any) {
	for _, arg := range args {
		result = append(result, arg)
	}
	return
}
