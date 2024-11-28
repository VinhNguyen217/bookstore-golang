package model

type Bill struct {
	ID          int    `gorm:"primaryKey"`
	UserId      int    `gorm:"user_id"`
	Receiver    string `gorm:"column:receiver"`
	Phone       string `gorm:"column:phone"`
	Address     string `gorm:"column:address"`
	Email       string `gorm:"column:email"`
	Note        string `gorm:"column:note"`
	Total       int    `gorm:"column:total"`
	Status      string `gorm:"column:status"` // Trạng thái đơn hàng
	Payment     string `gorm:"column:payment"`
	ConfirmDate string `gorm:"column:confirm_date"`
	CreatedDate string `gorm:"column:created_date"`
	UpdatedDate string `gorm:"column:updated_date"`
}

func (Bill) tableName() string {
	return "bills"
}
