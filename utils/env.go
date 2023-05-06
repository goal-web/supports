package utils

import (
	"github.com/goal-web/contracts"
	"os"
	"strings"
)

// LoadEnv 加载 .env 文件
func LoadEnv(envPath, sep string) (contracts.Fields, error) {
	envBytes, err := os.ReadFile(envPath)
	if err != nil {
		return nil, err
	}

	fields := make(contracts.Fields)
	for _, line := range strings.Split(string(envBytes), "\n") {
		if strings.HasPrefix(line, "#") { // 跳过注释
			continue
		}
		values := strings.Split(line, sep)
		if len(values) > 1 {
			fields[values[0]] = strings.Trim(strings.ReplaceAll(strings.Join(values[1:], sep), `"`, ""), "\r\t\v\x00")
		}
	}

	return fields, nil
}
