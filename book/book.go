package book

import "time"

type BookModel struct {
	ID          int
	Title       string
	Description string
	Price       int
	Discount    int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
