package main

import (
	"fmt"
)

func getMessageCost(message []string) []float64 {
	costArray := make([]float64, len(message))
	for i := 0; i < len(message); i++ {
		costMsg_i := float64(len(message[i])) * .01
		fmt.Println(costMsg_i)
		costArray[i] = costMsg_i
	}
	return costArray
}

func test(message []string) {
	costs := getMessageCost(message)
	fmt.Println("messages:")
	for i := 0; i < len(message); i++ {
		fmt.Printf(" - %v\n", message[i])
	}
	fmt.Println("costs: ")
	for j := 0; j < len(costs); j++ {
		fmt.Printf(" - %.2f\n", costs[j])
	}
	fmt.Println("=========END REPORT=========")
}

func main() {
	test([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	test([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	test([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})
}
