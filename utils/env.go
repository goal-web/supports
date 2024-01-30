package utils

import (
	"github.com/goal-web/contracts"
	"io/ioutil"
	"strings"
)

// LoadEnv 加载 .env 文件
func LoadEnv(envPath, sep string) (contracts.Fields, error) {
	envBytes, err := ioutil.ReadFile(envPath)
	if err != nil {
		return nil, err
	}

	// 将 \r\n 替换为 \n 以兼容 Windows 系统
	envContent := strings.ReplaceAll(string(envBytes), "\r\n", "\n")

	fields := make(contracts.Fields)
	for _, line := range strings.Split(envContent, "\n") {
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
