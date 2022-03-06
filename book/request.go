package book

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       interface{} `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	Rating      interface{} `json:"rating" binding:"required,number,min=1,max=10"`
	Discount    interface{} `json:"discount" binding:"required,number,max=100"`
}
