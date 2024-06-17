package entity

import (
	"time"

	"gorm.io/gorm"
)

type Discount struct {
	DiscountID      uint           `gorm:"primaryKey;autoIncrement" json:"discount_id"`
	DiscountCode    string         `gorm:"not null" json:"discount_code"`
	DiscountPrice   int            `json:"discount_price"`
	DiscountPercent int            `json:"discount_percent"`
	CreatedAt       time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
