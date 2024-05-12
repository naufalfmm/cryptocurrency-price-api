package dao

import "time"

type CoinHistory struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement;not null"`
	CoinID      uint64    `gorm:"not null"`
	LatestPrice float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	CreatedBy   string    `gorm:"not null"`

	Coin Coin `gorm:"foreignKey:ID;references:CoinID"`
}

func (CoinHistory) TableName() string {
	return "coin_histories"
}
