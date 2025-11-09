package repository

import "gorm.io/gorm"

// UserRepository 用户数据访问接口
type UserRepository interface {
	// TODO: 定义用户数据访问方法
}

// TaskRepository 任务数据访问接口
type TaskRepository interface {
	// TODO: 定义任务数据访问方法
}

// PlanRepository 计划数据访问接口
type PlanRepository interface {
	// TODO: 定义计划数据访问方法
}

// BaseRepository 提供通用功能
type BaseRepository struct {
	DB *gorm.DB
}

// NewBaseRepository 创建基础仓储
func NewBaseRepository(db *gorm.DB) BaseRepository {
	return BaseRepository{
		DB: db,
	}
}
