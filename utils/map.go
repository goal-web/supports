package utils

import (
	"fmt"
	"github.com/goal-web/contracts"
)

// GetMapKeys 获取 Fields 的所有 key
func GetMapKeys(data map[string]any) (keys []string) {
	for key := range data {
		keys = append(keys, key)
	}
	return
}

func Flatten(src, fields contracts.Fields, joiner string) {
	flattenMerge("", src, fields, joiner)
}

func flattenMerge(prefix string, src, fields contracts.Fields, joiner string) {
	for key, value := range fields {
		switch v := value.(type) {
		case map[string]any:
			flattenMerge(fmt.Sprintf("%s%s%s", prefix, key, joiner), src, v, joiner)
		default:
			src[prefix+key] = v
		}
	}
}
