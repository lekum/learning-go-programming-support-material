package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Name struct {
	First, Last string
}

type Book struct {
	Title       string    `json:"book_title"`
	PageCount   int       `json:"pages,string"`
	ISBN        string    `json:"-"`
	Authors     []Name    `json:"auths,omniempty"`
	Publisher   string    `json:",omniempty"`
	PublishDate time.Time `json:"pub_date"`
}

func main() {

	file, err := os.Open("books.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var books []Book
	dec := json.NewDecoder(file)
	if err := dec.Decode(&books); err != nil {
		fmt.Println(err)
	}
	fmt.Println(books)
}
