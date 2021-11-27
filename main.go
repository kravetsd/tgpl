package main

import (
	"os"

	"github.com/kravetsd/TheGoProgrammingLanguage/standardlib/io"
)

const (
	url string = "https://golang.org"
)

func main() {

	s, b := os.Args[1], os.Args[2]
	io.ReadFromString(s, b)
}
