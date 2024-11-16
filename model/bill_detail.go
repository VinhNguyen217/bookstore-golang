package model

import "time"

type BillDetail struct {
	ID        int `gorm:"primaryKey"`
	BillID    int `gorm:"column:bill_id"`
	BookID    int `gorm:"column:book_id"`
	Quantity  int `gorm:"column:quantity"`
	Price     int `gorm:"column:price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (BillDetail) TableName() string {
	return "bill_detail"
}
