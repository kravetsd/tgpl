package bytespkg

import (
	"bytes"
	"fmt"
	"os"
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

func Buffers() {
	var b bytes.Buffer // needs no initialization
	b.Write([]byte("Hello, ◺ this is string"))
	for i := 0; i < 5; i++ {
		fmt.Println(b.ReadByte())
	}

	fmt.Println(&b) //here we can see that someS first byte already read and triangle no more visible

	for i := 0; i < 3; i++ {

		rn, len, _ := b.ReadRune()
		// here we can see that Readrune reads erune by rune from the buffer.l
		fmt.Printf("%q, %d \n", rn, len)
	}
	fmt.Printf("%s\n", b.Bytes())

	// Now we are testing unread runes NOTE THE ONLY LAST RUNE CAN BE UNREADED:
	for i := 0; i < 10; i++ {

		b.UnreadRune()
		// here we can make unread operation for previously read runes:
	}
	fmt.Printf("%s\n", b.Bytes())

	//read a couple of bytes and see how it got changed: using Bytes function for buffer
	fmt.Println(b.Bytes())
	fmt.Println(b.ReadByte())
	fmt.Println(b.ReadByte())
	fmt.Println(b.Bytes())
	fmt.Printf("%s\n", b.Bytes())

	b.WriteString(" appended ◺ to the end of string")
	fmt.Printf("%s\n", b.Bytes())
	fmt.Println(b.ReadByte())
	fmt.Println(b.ReadByte())
	fmt.Printf("%s\n", b.Bytes())

	fmt.Println(b.ReadByte())
	fmt.Println(b.ReadByte())
	fmt.Printf("%s\n", b.Bytes())

	//since file is an interface io.Writer and stdout is a file too we can flush buffer into the file:

	b.WriteTo(os.Stdout)
	os.Stdout.WriteString("\n") //writing nrwline into stdout file
	//checking the lengt of the buffer:
	fmt.Println(b.Len())

}
