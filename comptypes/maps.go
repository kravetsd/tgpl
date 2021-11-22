package comptypes

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

var ages = map[string]int{"Mike": 34, "Sam": 7, "Helen": 33}

func sortmap(ages map[string]int) {
	fmt.Println(ages)
	var names []string
	var agear []int
	for key, val := range ages {
		names = append(names, key)
		agear = append(agear, val)
	}

	fmt.Println(agear, names)

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%v \t %d\n", name, ages[name])
	}
}

func equalmap(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, val := range m1 {
		if k, ok := m2[key]; ok != true || k != val {
			return false
		}
	}

	return true

}

func Dedup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(seen)

	}

}

// how we can make not comparable types to be amap keys:

func checkMapCompare() {
	var m = make(map[string]int)
	m = Add(m, []string{"hello", "world"})
	m = Add(m, []string{"hello", "world"})
	m = Add(m, []string{"hello", "world", "again"})

	fmt.Println(count(m, []string{"hello", "world"}))
	fmt.Println(count(m, []string{"hello", "world", "again"}))
	fmt.Println(m)
}

func Add(mp map[string]int, somelist []string) map[string]int {
	mp[k(somelist)]++
	return mp
}

func count(mp map[string]int, somelist []string) int {
	return mp[k(somelist)]
}

func k(slice []string) string { return fmt.Sprintf("%q", slice) }

// count UTF symbols in the input:

func Utflen() {
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	count := make(map[rune]int)

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		count[r]++
		utflen[n]++

	}

	fmt.Println(count, utflen, invalid)
	fmt.Printf("rune\tcount\n")
	for c, n := range count {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}

var m = make(map[string]map[string]bool)

func GraphInMao() {

	fmt.Println(m)
	AddEdge("NY", "LA")
	fmt.Println(m)
	fmt.Println(HasEdge("NY", "LA"))

}

func AddEdge(from, to string) {
	edges := m[from]
	if edges == nil {
		edges = make(map[string]bool)
		m[from] = edges
	}
	edges[to] = true
}

func HasEdge(from, to string) bool {
	return m[from][to]
}
