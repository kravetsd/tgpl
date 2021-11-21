package comptypes

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
