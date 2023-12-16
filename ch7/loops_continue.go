package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	for n := 2; n < (max + 1); n++ {
		if n == 2 {
			fmt.Println(n)
		}
		if n%2 == 0 {
			continue
		}
		continue__ := false
		for m := 3; m < (int(math.Sqrt(float64(n))) + 1); m++ {
			if n%m == 0 {
				continue__ = true
				break
			}
		}
		if continue__ {
			continue
		}
		fmt.Println(n)
	}
}

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("======================================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
