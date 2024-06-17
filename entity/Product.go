package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ProductID   uint           `gorm:"primaryKey;autoIncrement" json:"product_id"`
	ProductName string         `gorm:"not null" json:"product_name"`
	Price       float64        `gorm:"not null" json:"price"`
	Image       []byte         `gorm:"not null" json:"image"`
	Cost        float64        `gorm:"not null" json:"cost"`
	Reports     []Report       `gorm:"foreignKey:ProductID"` // One-to-many relationship
	CreatedAt   time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
