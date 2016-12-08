package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	// open input file
	file := "xml_goodread_42.xml"
	fi, err := os.Open(file)
	check(err)

	// defer to close file on exit and check err msg if any
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	// data := ReadAuthor(fi)
	dat, err := ioutil.ReadFile(file)
	check(err)
	fmt.Print(string(dat))
}
