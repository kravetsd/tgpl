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

func strtoslice(s string) {
	rn := []rune(s)
	fmt.Printf("%q", rn)

	runes := []rune{}
	for _, r := range s {
		runes = append(runes, r)
	}

	fmt.Printf("%q", runes)
}

func appendInt(x []int, y int) ([]int, int, int) {
	var z []int
	fmt.Println(cap(z))
	zlen := len(x) + 1
	fmt.Println(zlen, cap(x))
	if zlen <= cap(x) {
		z = append(x, y)
	} else {
		zcap := zlen * 2
		z = make([]int, zlen, zcap)
		fmt.Println(cap(z))
		copy(z, x)
	}

	z[len(x)] = y

	return z, len(z), cap(z)
}

func reversestr(st string) string {
	stb := []byte(st)
	for i, j := 0, len(stb)-1; i < j; i, j = i+1, j-1 {
		stb[i], stb[j] = stb[j], stb[i]
	}

	return string(stb)
}
