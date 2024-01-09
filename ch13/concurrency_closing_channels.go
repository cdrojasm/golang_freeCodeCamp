package main

import (
	"fmt"
	"time"
)

func countReports(numSentCh chan int) int {
	countTotalReportsSent := 0
	for {
		valueNumSent, ok := <-numSentCh
		if !ok {
			break
		}
		countTotalReportsSent += valueNumSent
	}

	return countTotalReportsSent
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Printf("sent batch of %v reports\n", numReports)
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

func test(numBatches int) {
	numSentCh := make(chan int)
	go sendReports(numBatches, numSentCh)
	fmt.Println("Start counting...")
	numReports := countReports(numSentCh)
	fmt.Printf("%v reports sent !\n", numReports)
	fmt.Println("==============================")
}

func main() {
	test(3)
	test(4)
	test(5)
	test(6)
}
