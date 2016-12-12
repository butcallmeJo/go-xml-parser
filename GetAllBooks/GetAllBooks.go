package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

// GetAllBooks uses http to get the author XML response from the goodreads API.
// authorID: string - ID associated to a specific author for the goodreads API.
// data: []byte - the http data response from the goodreads API.
func GetAllBooks(authorID string) []byte {
	// needed API key to access the GoodReads API - please have the key saved
	// as GOODREADS in your environment variables.
	apiKey := os.Getenv("GOODREADS")
	apiAuthorLink := "https://www.goodreads.com/author/show/" + authorID + ".xml?key=" + apiKey
	res, err := http.Get(apiAuthorLink)
	check(err)
	data, err := ioutil.ReadAll(res.Body)
	check(err)

	return data
}
