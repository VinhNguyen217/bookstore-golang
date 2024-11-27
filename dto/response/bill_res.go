package response

type BillRes struct {
	ID          int             `json:"id"`
	Receiver    string          `json:"receiver"`
	Phone       string          `json:"phone"`
	Address     string          `json:"address"`
	Email       string          `json:"email"`
	BillDetails []BillDetailRes `json:"billDetails"`
	Note        string          `json:"note"`
	Total       string          `json:"total"`
	Status      string          `json:"status"` // Trạng thái đơn hàng
	Payment     string          `json:"payment"`
	ConfirmDate string          `json:"confirmDate"` // Thời gian xác nhận đơn hàng
	CreatedDate string          `json:"createdDate"`
	UpdatedDate string          `json:"updatedDate"`
}
