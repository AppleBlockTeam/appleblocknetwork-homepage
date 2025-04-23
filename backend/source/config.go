package source

import (
	"fmt"
	"os"
	"strconv"

	"abnet_backend/source/logger"

	"github.com/joho/godotenv"
)

// ServerConfig 结构体定义服务器配置项
type ServerConfig struct {
	Host       string
	Port       int
	LG_BaseURL string
}

// Config 结构体定义配置项
type Config struct {
	Server ServerConfig
}

// 全局变量保存配置
var AppConfig *Config

// 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// 获取整数类型的环境变量，如果不存在或格式错误则返回默认值
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		logger.Warning(fmt.Sprintf("环境变量 %s 未填写，使用默认值: %d", key, defaultValue))
		return defaultValue
	}

	return value
}

// LoadConfig 加载配置
func LoadConfig() error {
	// 尝试加载.env文件
	loadEnvFile()

	serverHost := getEnv("SERVER_HOST", "0.0.0.0")
	serverPort := getEnvAsInt("SERVER_PORT", 8080)
	lgBaseURL := getEnv("LG_BASE_URL", "http://localhost:5000")

	AppConfig = &Config{
		Server: ServerConfig{
			Host:       serverHost,
			Port:       serverPort,
			LG_BaseURL: lgBaseURL,
		},
	}
	return nil
}

// loadEnvFile 尝试从当前目录和上级目录加载.env文件
func loadEnvFile() {
	// 尝试直接加载.env（当前目录）
	err := godotenv.Load()
	if err == nil {
		return
	}
}
