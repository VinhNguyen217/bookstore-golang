package model

import "time"

type Cart struct {
	ID        int `gorm:"primaryKey"`
	UserID    int `gorm:"column:user_id"`
	BookID    int `gorm:"column:book_id"`
	Quantity  int `gorm:"column:quantity"`
	Price     int `gorm:"column:price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Cart) TableName() string {
	return "cart"
}
