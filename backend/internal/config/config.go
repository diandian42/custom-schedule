package config

import (
	"fmt"
	"os"
)

// Config 应用配置
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

// ServerConfig 服务配置
type ServerConfig struct {
	Mode string
	Port string
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret     string
	ExpireHour int
}

// LoadConfig 加载配置（后续可扩展到读取文件或环境变量）
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Mode: getEnvOrDefault("APP_MODE", "debug"),
			Port: getEnvOrDefault("APP_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Driver: getEnvOrDefault("DB_DRIVER", "sqlite"),
			Host:   getEnvOrDefault("DB_HOST", "127.0.0.1"),
			Port:   getEnvOrDefault("DB_PORT", "3306"),
			User:   getEnvOrDefault("DB_USER", "root"),
			// Password 暂不强制提供，可在连接时校验
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   getEnvOrDefault("DB_NAME", "custom_schedule"),
		},
		JWT: JWTConfig{
			Secret:     getEnvOrDefault("JWT_SECRET", "replace-with-random-string"),
			ExpireHour: 24,
		},
	}

	// 预留扩展点：可以在此执行额外配置校验

	return cfg, nil
}

func getEnvOrDefault(key, def string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return def
}

// Addr 返回监听地址
func (s ServerConfig) Addr() string {
	return fmt.Sprintf(":%s", s.Port)
}
