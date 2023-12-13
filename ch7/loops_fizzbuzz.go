package main

import "fmt"

func fizzbuzz() {
	for i := 1; i < 101; i++ {
		isThreeMult := (i % 3) == 0
		isFiveMult := (i % 5) == 0
		if isThreeMult && isFiveMult {
			fmt.Println("fizzbuzz")
		} else if isThreeMult {
			fmt.Println("fizz")
		} else if isFiveMult {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz()
}
