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
	// here is a trick where we assign indices in a random order:
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RUR: "₽", RMB: "¥"}
	rubles := symbol[RUR]
	fmt.Println(RUR, rubles)
}

func SliceIntersection() {
	{
		s := ar[2:]
		q := ar[1:3]
		for _, v := range s {
			for _, val := range q {
				if v == val {
					fmt.Printf("%v is in both slices \n", v)
				}
			}
		}

	}

}

func reverse(ir []int) []int {
	for i, j := 0, len(ir)-1; i < j; i, j = i+1, j-1 {
		ir[i], ir[j] = ir[j], ir[i]
	}

	return ir
}

func reversep(ir *[5]int) []int {
	for i, j := 0, len(ir)-1; i < j; i, j = i+1, j-1 {
		ir[i], ir[j] = ir[j], ir[i]
	}

	return ir[:]
}
