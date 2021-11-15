package basictypes

import "fmt"

const GoUsage = `Go is a tool for managing Go source code.
Usage:
go command [arguments]
...`

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

}
