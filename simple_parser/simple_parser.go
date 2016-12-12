package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// func ReadAuthor(reader io.Reader) {

// 	return xmlAuthor
// }

func main() {

	type Author struct {
		XMLName xml.Name `xml:"author"`
		Name    string   `xml:"name"`
	}
	type Authors struct {
		XMLName xml.Name `xml:"authors"`
		Authors []Author `xml:"author"`
	}
	type Book struct {
		XMLName xml.Name  `xml:"book"`
		Book    []Authors `xml:"authors"`
	}
	type GoodreadsResponse struct {
		XMLName           xml.Name `xml:"GoodreadsResponse"`
		GoodreadsResponse []Book   `xml:"book"`
	}
	v := GoodreadsResponse{}

	// open input file
	// TODO: figure out the path situation for the xml
	file, err := filepath.Abs("xml_goodread_42.xml")
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

	err = xml.Unmarshal([]byte(dat), &v)
	check(err)

	fmt.Printf("name: %q\n", v.GoodreadsResponse[0].Book[0].Authors[0].Name)
}
