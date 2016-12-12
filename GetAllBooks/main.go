package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	AuthorID := XMLauthorParser()
	data := GetAllBooks(AuthorID)
	books := XMLbooksParser(data)
	fmt.Println(books)
}
