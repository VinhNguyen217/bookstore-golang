package request

type CartItemRequest struct {
	UserId   int `json:"userId"`
	BookId   int `json:"bookId"`
	Quantity int `json:"quantity"`
}
