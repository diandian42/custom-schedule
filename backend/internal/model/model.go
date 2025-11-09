package model

import "time"

// BaseModel 通用字段
type BaseModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TODO: 在设计数据库之后新增具体实体，例如 User、Task、Plan 等
