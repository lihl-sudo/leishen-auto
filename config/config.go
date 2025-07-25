package config

import (
	"fmt"
	"os"
)

// Config 配置结构体
type Config struct {
	AccountToken string
	Lang         string
}

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	accountToken := os.Getenv("TOKEN")
	if accountToken == "" {
		return nil, fmt.Errorf("TOKEN 环境变量未设置")
	}

	return &Config{
		AccountToken: accountToken,
		Lang:         "zh_CN",
	}, nil
}
