package comptypes

import (
	"fmt"
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
