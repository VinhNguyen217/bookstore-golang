package response

import "time"

type BookRes struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	Sold        int       `json:"sold"`
	Price       string    `json:"price"`
	PublishDate time.Time `json:"publishDate"`
	Description string    `json:"description"`
}
