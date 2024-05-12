package dao

import "time"

type Coin struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement;not null"`
	Code        string    `gorm:"not null"`
	CoincapID   string    `gorm:"not null"`
	Name        string    `gorm:"not null"`
	LatestPrice float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	CreatedBy   string    `gorm:"not null"`
	UpdatedBy   string    `gorm:"not null"`

	LatestPriceCurrency string `gorm:"-"`
}

func (Coin) TableName() string {
	return "coins"
}
