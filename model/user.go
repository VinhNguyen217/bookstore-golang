package model

import (
	"book-store/enum"
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Salt      string    `gorm:"column:salt"`
	Role      enum.Role `gorm:"column:role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "user"
}
