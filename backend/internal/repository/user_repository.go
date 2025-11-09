package repository

import "gorm.io/gorm"

type userRepository struct {
	BaseRepository
}

// NewUserRepository 创建用户仓储
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		BaseRepository: NewBaseRepository(db),
	}
}
