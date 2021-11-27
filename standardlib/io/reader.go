package io

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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

func ReadFromString(str, bufsize string) {

	// func main() {

	// 	s, b := os.Args[1], os.Args[2]
	// 	io.ReadFromString(s, b)
	// }

	st := strings.NewReader(str)
	size, err := strconv.ParseInt(bufsize, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	b := make([]byte, size)

	for {
		n, err := st.Read(b)

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		if n > 0 {
			fmt.Printf("%s\n", b)
		}
	}
}
