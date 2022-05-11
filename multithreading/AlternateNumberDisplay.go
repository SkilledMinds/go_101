package main

import (
	"fmt"
)

func main() {

	numFlag := make(chan int)
	done := make(chan bool)
	go oddNumberRoutine(numFlag, done)
	go evenNumberRoutine(numFlag, done)
	<-done
}

func oddNumberRoutine(numNeedToBePrinted chan int, printingCompleted chan bool) {

	for {
		v := <-numNeedToBePrinted
		if v >= 20 {
			printingCompleted <- true
			break
		}
		fmt.Println("Odd ==>", v)
		numNeedToBePrinted <- v + 1
	}

}

func evenNumberRoutine(numNeedToBePrinted chan int, printingCompleted chan bool) {
	numNeedToBePrinted <- 1
	for {
		v := <-numNeedToBePrinted
		if v > 20 {
			printingCompleted <- true
			break
		}
		fmt.Println("Even ==>", v)
		numNeedToBePrinted <- v + 1
	}

}
