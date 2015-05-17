package main

import (
	"fmt"
	"log"
)

func zipInt(list1 []int, list2 []int) []int {
	if len(list1) != len(list2) {
		log.Fatal("length mismatch")
	}

	var result []int

	for i, _ := range list1 {
		result = append(result, list1[i], list2[i])
	}

	return result
}

func main() {
	r := zipInt([]int{1, 2, 3}, []int{2, 3, 4})
	fmt.Println(r)
}
