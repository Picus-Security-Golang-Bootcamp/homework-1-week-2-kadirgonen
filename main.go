package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	Names string
}

func main() {

	bookJson := `[{"names":"Lord of The Rings"},{"names":"Harry Potter"},{"names":"Hunger Games"},{"names":"Treasure Island"},{"names":"Moonlight"},{"names":"Oliver Twist"}]`
	var books []Book
	json.Unmarshal([]byte(bookJson), &books)

	if os.Args[1] == "list" {

		ListBookNames(books)

	} else if os.Args[1] == "search" {

		var searchArg []string = os.Args[2:]
		fmt.Printf("Names : %+v", searchArg)

		args := os.Args[2:]
		for i, arg := range args {

			fmt.Printf("index: %d, value: %s\n", i, arg)
		}
		SearchBookName(books, strings.Join(searchArg, " "))

	}

}

func SearchBookName(books []Book, bookName string) {
	var bVal bool // default is false
	for i := range books {
		//fmt.Printf("Names : %+v", "bookName")
		if strings.EqualFold(books[i].Names, bookName) {
			bVal = true
			//fmt.Printf("Names : %+v", bookName)
		}
	}
	if bVal {
		fmt.Printf("Names : %+v", bookName)
	} else {
		fmt.Printf("Names : %+v", bookName+" Not Found")
	}

}

func ListBookNames(books []Book) {

	for _, s := range books {
		fmt.Println(s)
	}

}
