package request

type CartItemRequest struct {
	BookId   int `json:"bookId"`
	Quantity int `json:"quantity"`
}
