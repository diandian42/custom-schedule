package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	appConfig "custom-schedule/internal/config"
)

// InitDatabase 初始化数据库连接（后续可扩展更多驱动和配置）
func InitDatabase(cfg *appConfig.DatabaseConfig) (*gorm.DB, error) {
	switch cfg.Driver {
	case "mysql":
		return gorm.Open(mysql.Open(buildMySQLDSN(cfg)), &gorm.Config{})
	case "sqlite":
		return gorm.Open(sqlite.Open(buildSQLiteDSN(cfg)), &gorm.Config{})
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.Driver)
	}
}

func buildMySQLDSN(cfg *appConfig.DatabaseConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)
}

func buildSQLiteDSN(cfg *appConfig.DatabaseConfig) string {
	if cfg.DBName == "" {
		return "custom_schedule.db"
	}
	return fmt.Sprintf("%s.db", cfg.DBName)
}
