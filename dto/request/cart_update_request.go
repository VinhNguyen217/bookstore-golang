package request

type CartItemUpdateRequest struct {
	CartId   int `json:"cartId"`
	Quantity int `json:"quantity"`
}
