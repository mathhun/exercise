package main

import (
	"fmt"
	"os"
	"strconv"
)

func sumWithFor(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func sumWithWhile(nums []int) int {
	i := 0
	sum := 0
	for i < len(nums) {
		sum += nums[i]
		i += 1
	}
	return sum
}

func sumWithRecursion(nums []int, accum int) int {
	if len(nums) == 0 {
		return accum
	}

	return sumWithRecursion(nums[1:], accum+nums[0])
}

func solve(nums []int) {
	fmt.Println(sumWithFor(nums))
	fmt.Println(sumWithWhile(nums))
	fmt.Println(sumWithRecursion(nums, 0))
}

func main() {
	var nums []int
	for _, n := range os.Args[1:] {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}
	solve(nums)
}
