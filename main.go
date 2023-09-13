package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func loadJSON() []book {
	file, err := os.Open("books.json")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var books []book

	json.Unmarshal(byteValue, &books)

	return books
}

var books = loadJSON()

func searchBookID(id int) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found!")
}

func getAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)

	c.IndentedJSON(http.StatusCreated, newBook)

	createJSON(books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	book, err := searchBookID(i)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book Not Found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getAllBooks)
	router.POST("/add", addBook)
	router.GET("/books/:id", getBookByID)
	router.Run("localhost:8080")
}
