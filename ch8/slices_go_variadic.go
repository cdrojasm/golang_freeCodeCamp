package main

import (
	"fmt"
)

func sum(nums ...float64) float64 {
	sumCost := 0.0
	for i := 0; i < len(nums); i++ {
		sumCost += nums[i]
	}
	return sumCost
}

func test(num ...float64) {
	total := sum(num...)
	fmt.Printf("Summing %v costs ...\n", len(num))
	fmt.Printf("Bill for the month %.2f\n", total)
	fmt.Printf("========END OF REPORT============")

}

func main() {
	test(1.0, 2.0, 3.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)
}
