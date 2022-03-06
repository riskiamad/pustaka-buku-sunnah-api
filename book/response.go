package book

type BookResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
	Discount      int    `json:"discount"`
	Rating        int    `json:"rating"`
	DiscountPrice int    `json:"discountPrice"`
}