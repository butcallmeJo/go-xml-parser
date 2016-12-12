package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// XMLauthorParser parses a GoodreadsResponse Book XML file in
// the bin dir and returns the author (ID) of the book.
func XMLauthorParser() string {
	// Author is a sub-type of Authors. Contains the Name, web Link and ID
	// of the author in question.
	type Author struct {
		XMLName xml.Name `xml:"author"`
		Name    string   `xml:"name"`
		Link    string   `xml:"link"`
		ID      string   `xml:"id"`
	}
	// Authors is a sub-type of Book. Contains all the potential authors
	// of the book.
	type Authors struct {
		XMLName xml.Name `xml:"authors"`
		Authors []Author `xml:"author"`
	}
	// Book is a sub-type of the main Response. Contains the authors of the
	// book.
	type Book struct {
		XMLName xml.Name  `xml:"book"`
		Book    []Authors `xml:"authors"`
	}
	// GoodreadsResponse is the main response. Contains the book.
	type GoodreadsResponse struct {
		XMLName           xml.Name `xml:"GoodreadsResponse"`
		GoodreadsResponse Book     `xml:"book"`
	}
	v := GoodreadsResponse{}

	// open input file
	// TODO in future implementation: Organize path and dir for XML files
	file, err := filepath.Abs("xml_goodread_42.xml")
	check(err)
	fi, err := os.Open(file)
	check(err)

	// defer to close file on exit and check err msg if any
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	dat, err := ioutil.ReadAll(fi)
	check(err)

	// Associating the XML data to the structs created above.
	err = xml.Unmarshal([]byte(dat), &v)
	check(err)

	author := v.GoodreadsResponse.Book[0].Authors[0].Name
	link := v.GoodreadsResponse.Book[0].Authors[0].Link
	id := v.GoodreadsResponse.Book[0].Authors[0].ID
	fmt.Printf("author name: %s\n", author)
	fmt.Printf("author link: %s\n", link)
	fmt.Printf("author id: %s\n", id)
	return id
}

// XMLbooksParser parses a GoodreadsResponse Author XML file taken from
// the web and returns a list of books written by the same author.
func XMLbooksParser(data []byte) []string {
	// BookB is a sub-type of BooksB. Contains the Title of the Book.
	type BookB struct {
		XMLName xml.Name `xml:"book"`
		Title   string   `xml:"title"`
	}
	// BooksB is a sub-type of AuthorB. Contains all the Books written by the
	// author.
	type BooksB struct {
		XMLName xml.Name `xml:"books"`
		BookB   []BookB  `xml:"book"`
	}
	// AuthorB is a sub-type of the main Response. Contains all the information
	// of the author: Name, web Link, ID, list of Books.
	type AuthorB struct {
		XMLName xml.Name `xml:"author"`
		Name    string   `xml:"name"`
		Link    string   `xml:"link"`
		ID      string   `xml:"id"`
		BooksB  []BooksB `xml:"books"`
	}
	// GoodreadsResponseB is the main Response. Contains the Author.
	type GoodreadsResponseB struct {
		XMLName            xml.Name  `xml:"GoodreadsResponse"`
		GoodreadsResponseB []AuthorB `xml:"author"`
	}
	b := GoodreadsResponseB{}

	// Associating the XML data to the structs created above.
	err := xml.Unmarshal([]byte(data), &b)
	check(err)

	books := b.GoodreadsResponseB[0].BooksB[0].BookB
	booksSlice := make([]string, len(books))
	// Associating the books multi-level struct to a simpler slice of books
	for key, value := range books {
		booksSlice[key] = value.Title
	}
	return booksSlice
}
