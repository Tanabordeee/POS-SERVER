package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	EmployeeID uint           `gorm:"primaryKey;autoIncrement" json:"employee_id"`
	Name       string         `gorm:"not null" json:"name"`
	Role       string         `gorm:"not null" json:"role"`
	Salary     float64        `gorm:"not null" json:"salary"`
	CreatedAt  time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Auth       Auth           `gorm:"foreignKey:EmployeeID"`
}
