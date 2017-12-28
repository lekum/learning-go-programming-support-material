package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type Name struct {
	First, Last string
}

type Book struct {
	Title       string
	PageCount   int
	ISBN        string
	Authors     []Name
	Publisher   string
	PublishDate time.Time
}

func main() {

	books := []Book{
		Book{
			Title:     "Leaning Go",
			PageCount: 375, ISBN: "9781784395438",
			Authors:     []Name{{"Vladimir", "Vivien"}},
			Publisher:   "Packt",
			PublishDate: time.Date(2016, time.July, 0, 0, 0, 0, 0, time.UTC),
		},
		Book{
			Title:       "The Go Programming Language",
			PageCount:   380,
			ISBN:        "9780134190440",
			Authors:     []Name{{"Alan", "Donavan"}, {"Brian", "Kernighan"}},
			Publisher:   "Addison-Wesley",
			PublishDate: time.Date(2015, time.October, 26, 0, 0, 0, 0, time.UTC),
		},
	}

	file, err := os.Create("book.dat")
	if err != nil {
		fmt.Println(err)
		return
	}

	enc := gob.NewEncoder(file)
	if err := enc.Encode(books); err != nil {
		fmt.Println(err)
	}
}
