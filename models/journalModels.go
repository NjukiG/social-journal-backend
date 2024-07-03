package models

import "gorm.io/gorm"

type Role string

const (
	RoleMember Role = "Author"
	RoleAdmin  Role = "Admin"
)

// User Object/Model
type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Role      Role   `gorm:"not null"`
}
