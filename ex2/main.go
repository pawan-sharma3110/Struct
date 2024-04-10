package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type BookStore struct {
	Books []Books
}
type Books struct {
	Title   string
	Subject string
	Author  string
	Price   int
}

var AvliableBooks = BookStore{
	Books: []Books{
		{Title: "Rise to World",
			Subject: "Entertanment",
			Author:  "Arwind",
			Price:   250,
		},
		{Title: "Air",
			Subject: "Nature",
			Author:  "Asha",
			Price:   450,
		},
		{Title: "One Day",
			Subject: "Comix",
			Author:  "Vandhana",
			Price:   150,
		},
	},
}

func main() {
	data, err := json.MarshalIndent(AvliableBooks, " ", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))

}
