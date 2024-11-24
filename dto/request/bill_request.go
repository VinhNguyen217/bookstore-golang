package request

type BillRequest struct {
	CartIds  []int  `json:"cartIds"`
	Receiver string `gorm:"column:receiver"`
	Phone    string `gorm:"column:phone"`
	Address  string `gorm:"column:address"`
	Email    string `gorm:"column:email"`
	Note     string `gorm:"column:note"`
}
