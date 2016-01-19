package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"a", "d", "b", "c"}
	sort.Strings(strs)
	fmt.Println("String sort", strs)

	ints := []int{2, 1, 5, 3, 4}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)

}
