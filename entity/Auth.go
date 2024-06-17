package entity

type Auth struct {
	AuthID     uint   `gorm:"primaryKey;autoIncrement" json:"auth_id"`
	EmployeeID uint   `gorm:"not null" json:"employee_id"`
	Username   string `gorm:"unique;not null" json:"username"`
	Password   string `gorm:"not null" json:"password"`
}
