package book

import "time"

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      string
	Discount    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
