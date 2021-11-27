package main

import (
	"os"

	"github.com/kravetsd/TheGoProgrammingLanguage/standardlib/io"
)

const (
	url string = "https://golang.org"
)

func main() {

	path := os.Args[1]
	io.ReaderFormFile(path)
}
