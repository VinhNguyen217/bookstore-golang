package response

type BillDetailRes struct {
	ID        int    `json:"id"`
	BookID    int    `json:"bookId"`
	BookName  string `json:"bookName"`
	Quantity  int    `json:"quantity"`
	Price     string `json:"price"`
	UnitPrice string `json:"unitPrice"`
}
