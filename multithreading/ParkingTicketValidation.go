package main

//create a channel to prevent the main program from exiting before the done signal is received

import (
	"fmt"
	"time"
)

var ticketIDs = []int{3, 23, 4, 12, 78, 26, 103}
var validTicketIDs = []int{56, 73, 10, 13, 73, 13, 113, 78}

func main() {

	// channel first created
	ch := make(chan int)

	// channel second created
	done := make(chan bool)

	// let's execute ticket validation process through two separate goroutines ( two threads )
	go displayParkingTicketID(ch)
	go verifyParkingTicketID(ch, done)

	//This will prevent the program from exiting till a value is sent over the "done" channel, value doesn't matter
	// It is similar to Join() of Java
	<-done
	fmt.Println("Today's task is done. Let's party now !")
}

func displayParkingTicketID(ch chan int) {

	for i := range ticketIDs {
		// will wait for 2 second for next person in parking
		t := time.NewTimer(time.Second * 2)
		<-t.C
		fmt.Println("Sending ticketID to channel. ticketID : ", ticketIDs[i])
		//this goroutine will wait till another goroutine received the value
		ch <- ticketIDs[i]
	}
}

func verifyParkingTicketID(ch chan int, done chan bool) {
	ticketValidated := 0
	for {
		//this goroutine will wait till the channel ch received a value
		v := <-ch
		if isTicketExists(v) {
			fmt.Println(v, " is valid parking ticket")
		} else {
			fmt.Println(v, " is not valid parking ticket")
		}
		ticketValidated++
		if ticketValidated == len(ticketIDs) {
			break
		}
	}
	done <- true
}

func isTicketExists(ticket int) bool {
	for validTicket := range validTicketIDs {
		if ticket == validTicket {
			return true
		}
	}
	return false
}
