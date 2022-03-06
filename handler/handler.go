package handler

import (
	"fmt"
	"net/http"
	"pustaka-buku-sunnah-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON((&bookRequest))
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, Condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON((&bookRequest))
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, Condition : %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	id := c.Param("id")
	intid, _ := strconv.Atoi(id)

	book, err := h.bookService.Update(intid, bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		discountPrice := b.Price - (b.Price * b.Discount / 100)
		bookResponse := book.BookResponse{
			ID:            b.ID,
			Title:         b.Title,
			Description:   b.Description,
			Price:         b.Price,
			Discount:      b.Discount,
			Rating:        b.Rating,
			DiscountPrice: discountPrice,
		}

		booksResponse = append(booksResponse, bookResponse)
	}

	data := gin.H{
		"data": booksResponse,
	}

	c.JSON(http.StatusOK, data)
}

func (h *bookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)

	b, err := h.bookService.FindByID(intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(b)

	data := gin.H{
		"data": bookResponse,
	}

	c.JSON(http.StatusOK, data)
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)

	b, err := h.bookService.Delete(intId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func convertToBookResponse(b book.BookModel) book.BookResponse {
	discountPrice := b.Price - (b.Price * b.Discount / 100)
	return book.BookResponse{
		ID:            b.ID,
		Title:         b.Title,
		Description:   b.Description,
		Price:         b.Price,
		Discount:      b.Discount,
		Rating:        b.Rating,
		DiscountPrice: discountPrice,
	}
}
