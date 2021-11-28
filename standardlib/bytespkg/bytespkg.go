package bytespkg

import (
	"fmt"
	"unicode/utf8"
)

func BasicBytes() {
	s := "this is a string"
	bbs := []byte(s)
	bs := []byte{71, 111}
	fmt.Printf("%s\n", bs)
	fmt.Printf("%s\n", bbs)
	fmt.Printf("%d\n", bbs)

	// some symbols and bytes representing it:
	nonascii := []byte("◺")
	fmt.Printf("%s\n", nonascii)
	fmt.Printf("%d\n", nonascii)
}

func BasicEncoding() {
	bs := []byte("◺")
	s := string(bs)

	fmt.Println(s)
	fmt.Printf("%d\n", bs)

	fmt.Println(utf8.RuneCountInString(s))

	for i, val := range "Hello, ◺ this is string" {
		fmt.Printf("i: %v\t val: %d\t letter: %q\t len: %d\t %c\n", i, val, val, utf8.RuneLen(val), val)
	}
}
