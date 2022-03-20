package main

import (
	"log"
	"net/http"
	"os"
	"pustaka-buku-sunnah-api/book"
	"pustaka-buku-sunnah-api/handler"
	"regexp"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const projectDirName = "pustaka-buku-sunnah"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	loadEnv()

	host := os.Getenv("HOST")
	userName := os.Getenv("USERN")
	pass := os.Getenv("PASS")
	database := os.Getenv("DB")
	dataSource := "host=" + host + " user=" + userName + " password=" + pass + " database=" + database + " port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&book.BookModel{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/v1/books")
	})
	v1 := router.Group("/v1")
	v1.POST("/books", (*bookHandler).CreateBook)
	v1.GET("/books", (*bookHandler).GetBooks)
	v1.GET("/books/:id", (*bookHandler).GetBook)
	v1.PUT("/books/:id", (*bookHandler).UpdateBook)
	v1.DELETE("/books/:id", (*bookHandler).DeleteBook)

	router.Run(":8080")
}
