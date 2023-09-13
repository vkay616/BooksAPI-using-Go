package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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

func getAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getAllBooks)
	router.Run("localhost:8080")
}
