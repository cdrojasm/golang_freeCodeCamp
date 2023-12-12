package main

import (
	"fmt"
)

func bulkSend(numMessage int) float64 {
	//?
	cumulativeCost := 0.0
	baseTax := 1.0
	for fee := 0.0; fee < float64(numMessage)*0.01; fee += 0.01 {
		cumulativeCost += baseTax + fee
	}
	return cumulativeCost
}

func test(numMessage int) {
	fmt.Printf("Sending %v message\n", numMessage)
	cost := bulkSend(numMessage)
	fmt.Printf("Bulk send complete! Cost = %.2f\n", cost)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
	test(40)
	test(50)
}
