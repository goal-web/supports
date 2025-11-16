package utils

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var src = rand.NewSource(time.Now().UnixNano())

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

// RandStr 生成随机字符串
func RandStr(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			sb.WriteByte(letters[idx])
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return sb.String()
}

func RandIntArray(min, max, num int) []int {
	if min > max {
		min, max = max, min
	}
	results := make([]int, 0)
	randSource := rand.NewSource(time.Now().UnixNano())
	if num > max-min { // 有重复
		for i := 0; i < num; i++ {
			results = append(results, RandIntBySeed(randSource, min, max))
		}
	} else {
		for len(results) < num {
			i := RandIntBySeed(randSource, min, max)
			if IsNotIn(i, results) {
				results = append(results, i)
			}
		}
	}
	return results
}

func RandInt(min, max int) int {
	return RandIntBySeed(src, min, max)
}

func RandIntBySeed(source rand.Source, min, max int) int {
	return rand.New(source).Intn(int(math.Abs(float64(min))+math.Abs(float64(max)))+1) + min
}

// RandomStringBySeed 根据种子生成确定性的随机字符串
// seed: 随机种子（相同种子生成相同结果）
// length: 字符串长度
// charset: 可选字符集（默认包含大小写字母和数字）
func RandomStringBySeed(seed int64, length int, charset ...string) string {
	// 初始化随机源（确保结果确定性）
	r := rand.New(rand.NewSource(seed))

	// 定义默认字符集（大小写字母 + 数字）
	defaultCharset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	// 使用自定义字符集（如果提供）
	targetCharset := defaultCharset
	if len(charset) > 0 && charset[0] != "" {
		targetCharset = charset[0]
	}

	// 将字符集转换为字符切片
	characters := []rune(targetCharset)
	charsetLength := len(characters)

	// 生成随机字符串
	result := make([]rune, length)
	for i := range result {
		result[i] = characters[r.Intn(charsetLength)]
	}

	return string(result)
}
