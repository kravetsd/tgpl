package bytespkg

import (
	"bytes"
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

	a := '`'
	fmt.Printf("%q\n", a)
	for i, val := range "Hello, ◺ this is string" {
		fmt.Printf("i: %v\t val: %d\t letter: %q\t len: %d\t %c\n", i, val, val, utf8.RuneLen(val), val)
	}

	fmt.Println("Triangle is in the", bytes.IndexRune([]byte("Hello, ◺ this is string"), '◺'), "place")

	//fmt.Printf("%#v", bytes.Replace([]byte("Hello, ◺ this is string"), []byte("◺"), []byte("WTF"), -1))

	runes := bytes.Runes([]byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0x57, 0x54, 0x46, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x69, 0x73, 0x20, 0x73,
		0x74, 0x72, 0x69, 0x6e, 0x67})

	fmt.Println(string(runes))
}
