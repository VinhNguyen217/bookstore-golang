package request

type CartUpdateRequest struct {
	id       int `json:"id"`
	quantity int `json:"quantity"`
}
