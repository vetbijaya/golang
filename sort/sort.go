package main

import (
	"fmt"
	"sort"
)

func main() {
	ai := []int{3, 2, 5, 7, 6, 8, 9, 1}
	as := []string{"Nepal", "Dr.", "Tara", "Miss"}
	// print the values before and after sorting
	fmt.Println(ai)
	sort.Ints(ai) // sorting the integer elements
	fmt.Println(ai)

	fmt.Println("-------------")

	fmt.Println(as)
	sort.Strings(as) // sorting the strings elements
	fmt.Println(as)
}
