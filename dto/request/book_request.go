package request

type BookRequest struct {
	Name        string `json:"name"`
	Quantity    int    `json:"quantity"`
	Price       int    `json:"price"`
	PublishDate string `json:"publishDate"`
	Description string `json:"description"`
}
