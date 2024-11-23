package model

import (
	"time"
)

type Book struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"column:name"`
	Quantity    int       `gorm:"column:quantity"`
	Sold        int       `gorm:"column:sold"`
	Price       int       `gorm:"column:price"`
	PublishDate time.Time `gorm:"column:publish_date"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"<-:create"` // allow read and create, but don't update
	UpdatedAt   time.Time
}

func (Book) TableName() string {
	return "book"
}
