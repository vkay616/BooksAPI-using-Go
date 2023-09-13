package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  string `json:"price"`
}

func scrape() {
	c := colly.NewCollector()

	i := 1

	var books []book

	c.OnHTML("div.col-sm-20 div.card.align-items-center div.card-body.position-relative p.card-text.text-center", func(h *colly.HTMLElement) {
		book := book{
			ID:     i,
			Title:  h.ChildText("span.booktitle"),
			Author: h.ChildText("span.author"),
			Price:  h.ChildText("span.actualprice"),
		}

		i += 1

		// fmt.Println(h.ChildText("span.booktitle"))
		// fmt.Println(h.ChildText("span.author"))
		// fmt.Println(h.ChildText("span.actualprice"))

		books = append(books, book)
	})

	c.Visit("https://www.bookswagon.com/promo-best-seller/best-seller/03AC998EBDC2")

	// fmt.Println(books)

	content, err := json.Marshal(books)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("books.json", content, 0644)
}
