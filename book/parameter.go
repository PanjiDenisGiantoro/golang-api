package book

import "encoding/json"

type BookInput struct {
	Title    string      `json:"title,omitempty" binding:"required"`
	Price    json.Number `json:"price" binding:"required,number"`
	SubTitle string      `json:"sub_title" json:"sub_title,omitempty"`
}
