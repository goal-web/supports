package utils

// GetMapKeys 获取 Fields 的所有 key
func GetMapKeys(data map[string]interface{}) (keys []string) {
	for key, _ := range data {
		keys = append(keys, key)
	}
	return
}
