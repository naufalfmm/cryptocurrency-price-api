package dao

import (
	"time"

	"gorm.io/gorm"
)

type UserCoin struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement;not null"`
	UserID      uint64         `gorm:"not null"`
	CoinID      uint64         `gorm:"not null"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"null"`
	CreatedBy   string         `gorm:"not null"`
	UpdatedBy   string         `gorm:"not null"`
	DeletedBy   *string        `gorm:"null"`
	DeletedUnix int            `gorm:"not null"`

	Coin Coin `gorm:"foreignKey:ID;references:CoinID"`
	User User `gorm:"foreignKey:ID;references:UserID"`
}

func (UserCoin) TableName() string {
	return "user_coins"
}
