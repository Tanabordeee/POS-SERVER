package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    uint           `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username  string         `gorm:"unique;not null" json:"username"`
	Password  string         `gorm:"not null" json:"password"`
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
