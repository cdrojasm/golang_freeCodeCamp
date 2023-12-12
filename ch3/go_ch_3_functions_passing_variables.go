package main 

import "fmt"

func main (){
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("You've sent ", sendsSoFar, " messages")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	return sendsSoFar + sendsToAdd
}