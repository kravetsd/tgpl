package io

import (
	"fmt"
	"io"
	"os"
)

func ReaderFormFile(fpath string) {
	// func main() {

	// 	path := os.Args[1]
	// 	io.ReaderFormFile(path)
	// }

	f, err := os.Open(fpath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b := make([]byte, 1024)

	for {
		_, err := f.Read(b)
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", b)
	}
}
