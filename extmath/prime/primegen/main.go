package main

import (
	"fmt"
	"github.com/larsmans/algo/extmath/prime"
)

func main() {
	var sieve prime.Sieve32
	pr := make([]uint32, 10)
	for {
		for _, p := range sieve.Next(pr) {
			fmt.Println(p)
		}
	}
}
