package basictypes

import (
	"bytes"
	"fmt"
)

const GoUsage = `Go is a tool for managing Go source code.
Usage:
go command [arguments] 漢字
...`

var r, u, n = 'r', 'u', 'n'
var b byte = 12

func Strs() {
	a := "this is a string"
	fmt.Println("length of the string 'a' is : ", len(a))
	fmt.Println("0, 1-st and 15-th bytes of the string", a[0], a[2], a[15])
	fmt.Println("bytes if the string from 0 to 5 resulting in substring: ", a[0:6], a[0:6][2])

	fmt.Println("string concatenation :", a+a[:6])

	s := "left foot"
	t := s
	s += ", right foot"
	fmt.Println("Strings are immutable: ", s, t)

	// s[0] = "T" This is a compile error

	p := s[:]
	z := s[:]

	fmt.Println("substrings are memory cheap", p, z)

	fmt.Println("this is a raw string example : ", GoUsage)

	fmt.Println("\054")

	// bring some mutability into string
	st := "this is a string"
	b := []byte(st)
	fmt.Println(b)
	b[0] = 12

	fmt.Println(string(b)) //resulting string

	fmt.Println("converting a slice of integers to string: ", IntsToString([]int{1, 2, 3, 4}))

}

func Somestrconversion() {
	fmt.Println("Hello, playground")
	fmt.Printf("%T\n", GoUsage)
	fmt.Printf("%T\n", r)
	fmt.Printf("%T\n", b)
	c := byte(r)
	fmt.Printf("%v\n", string(c))

	stringa := string(r)
	fmt.Println(stringa)

	newstr := []rune(GoUsage)
	fmt.Println(newstr)
	fmt.Println([]byte(GoUsage))
	fmt.Println(string(newstr))

	n := "\n"
	fmt.Printf("%v,%v,%[2]T,%v,%[3]T", len(n), '\n', byte('\n')) //a trick with stri gformatting

}

func basepath(path string) {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			path = path[i+1:]
			fmt.Println(path)
			break
		}
	}

	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			path = path[:i]
			break
		}
	}

	fmt.Println(path)
}

// converting slice of int to string using bytes package:
func IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, val := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", val)
	}

	buf.WriteByte(']')
	return buf.String()
}
