package models

import "gorm.io/gorm"

type Role string

const (
	RoleAuthor Role = "Author"
	RoleAdmin  Role = "Admin"
)

// User Object/Model
type User struct {
	gorm.Model
	FirstName  string     `gorm:"not null"`
	LastName   string     `gorm:"not null"`
	Email      string     `gorm:"unique;not null"`
	Password   string     `gorm:"not null"`
	Role       Role       `gorm:"not null"`
	Categories []Category `gorm:"foreignKey:UserID"`
	Journals   []Journal  `gorm:"foreignKey:UserID"`
}

// Category struct/ object
type Category struct {
	gorm.Model
	Title    string    `gorm:"not null"`
	Journals []Journal `gorm:"foreignKey:CategoryID"`
	UserID   uint
}

type Journal struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Content    string `gorm:"not null"`
	ImageURL   string
	UserID     uint
	CategoryID uint
}
