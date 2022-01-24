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
	return rand.New(src).Intn(int(math.Abs(float64(min))+math.Abs(float64(max)))) + min
}

func RandIntBySeed(source rand.Source, min, max int) int {
	return rand.New(source).Intn(int(math.Abs(float64(min))+math.Abs(float64(max)))) + min
}
