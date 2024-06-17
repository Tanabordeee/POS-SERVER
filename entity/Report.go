package entity

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	ReportID     uint           `gorm:"primaryKey;autoIncrement" json:"report_id"`
	ReportDate   time.Time      `gorm:"not null" json:"report_date"`
	QuantitySold int            `gorm:"not null" json:"quantity_sold"`
	UpdatedAt    time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ProductID    uint           `gorm:"not null" json:"product_id"`
	Product      Product        `json:"product"`
}
