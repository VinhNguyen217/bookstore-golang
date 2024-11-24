package response

import "time"

type BillRes struct {
	ID          int       `json:"id"`
	Receiver    string    `json:"receiver"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	Email       string    `json:"email"`
	Note        string    `json:"note"`
	Total       string    `json:"total"`
	Status      string    `json:"status"` // Trạng thái đơn hàng
	Payment     string    `json:"payment"`
	ConfirmDate time.Time `json:"confirmDate"`
}
