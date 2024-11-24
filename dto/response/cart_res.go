package response

type CartRes struct {
	ID       int    `json:"id"`
	BookID   int    `json:"bookId"`
	Quantity int    `json:"quantity"`
	Price    string `json:"price"`
}
