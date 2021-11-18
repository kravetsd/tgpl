package comptypes

import (
	"fmt"
)

var ar = [5]int{1, 2, 3, 4, 5}

func Arrays() {
	fmt.Println("this is composite types package")
	fmt.Println("Hello, playground")
	for i, v := range ar {
		fmt.Printf("%d , %v \n", i, v)
	}

	for _, v := range ar {
		fmt.Printf("%v\n", v)
	}

	// elipsis specification for array:

	a := [...]int{1, 2, 3, 4}

	fmt.Printf("%T", a)

}

type Currency int8

const (
	USD Currency = iota
	GBP
	EUR
	RMB
	RUR
)

func Cury() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RUR: "₽", RMB: "¥"}
	dollar := symbol[USD]
	fmt.Println(dollar)
}
