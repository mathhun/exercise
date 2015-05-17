package main

import (
	"fmt"
	"log"
	"math/big"
)

var (
	zero = *big.NewInt(0)
	one  = *big.NewInt(1)
)

func fib(n int) []big.Int {
	if n < 1 {
		log.Fatal("n must be >= 1")
	}

	switch n {
	case 1:
		return []big.Int{zero}
	case 2:
		return []big.Int{one}
	default:
		return fibRecur(n-2, []big.Int{zero, one})
	}
}

func fibRecur(n int, list []big.Int) []big.Int {
	if n == 0 {
		return list
	}

	l := len(list)
	i := new(big.Int).Add(&list[l-2], &list[l-1])
	list = append(list, *i)
	return fibRecur(n-1, list)
}

func main() {
	r := fib(100)
	for i, n := range r {
		fmt.Printf("%4d: %v\n", i+1, &n)
	}
}
