package dao

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;not null"`
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	CreatedBy string    `gorm:"not null"`
	UpdatedBy string    `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}

type UserSignIn struct {
	Token string
	User  User
}
