package comptypes

import (
	"fmt"
	"strings"
	"time"
)

func Equality(l, k []int) bool {
	fmt.Println("comparing the lenght of both slices:")
	for _, v := range strings.Repeat(".", 10) {
		fmt.Printf("%v", v)
		time.Sleep(3 * time.Second)
	}

	if len(l) != len(k) {
		return false
	}

	for i := 0; i < len(l); i++ {
		if l[i] != k[i] {
			return false
		}
	}

	return true
}
