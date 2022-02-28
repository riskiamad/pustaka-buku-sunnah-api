package main

import (
	"log"
	"net/http"
	"pustaka-buku-sunnah-api/controller"
	book "pustaka-buku-sunnah-api/models"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=ec2-54-156-110-139.compute-1.amazonaws.com user=bigtiabtdobzpl password=64f926ac6c655587a230c09c693cdef90b1ef8125664a5c5d60c3968bc6644ae dbname=dfm5hvm7itml0i sslmode=require TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&book.BookModel{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := controller.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.POST("/books", (*bookHandler).CreateBook)
	v1.GET("/books", (*bookHandler).GetBooks)
	v1.GET("/books/:id", (*bookHandler).GetBook)
	v1.PUT("/books/:id", (*bookHandler).UpdateBook)
	v1.DELETE("/books/:id", (*bookHandler).DeleteBook)
	v1.GET("/mark", func(c *gin.Context) {
		c.String(http.StatusOK, string(blackfriday.Run([]byte("**hi!**"))))
	})

	router.Run()
}
