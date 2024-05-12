package dao

type Option struct {
	ID    uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Key   string `gorm:"not null"`
	Value string `gorm:"not null"`
}

func (Option) TableName() string {
	return "options"
}
