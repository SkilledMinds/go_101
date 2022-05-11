package main

// Implementation : custom implementation of Stack

// What all are covered here : Generics, Panic() , Recover() , Error handling

// TODO : Change to thread safe Stack
// TODO : Implement dynamic capacity increase once threshold reach
// TODO : Add Load factor percentage to support dynamic capacity

import (
	"fmt"
	"log"
	"strconv"
)

import (
	"errors"
)

var stackElements [STACK_SIZE]any
var currentIndex = 0

const STACK_SIZE = 10

func main() {

	var j int

	fmt.Println(" ******  PUSH elements to stack   ******")
	for j = 0; j < 11; j++ {
		objTostk := "A:" + strconv.Itoa(j)
		push(objTostk)
	}

	fmt.Println(" ******  POP elements from stack   ******")

	for j = 0; j < 11; j++ {
		object, errorInPop := pop()
		if errorInPop != nil {
			fmt.Println("No more element for POP")
		} else {
			fmt.Println("POPed => ", object)
		}

	}

}

// PUSH generic objects
func push[T any](objectToStk T) {

	if currentIndex >= STACK_SIZE {
		defer func() {
			if err := recover(); err == nil {
				log.Println("stack overflow occurred", err)
			}
			currentIndex--
		}()
		// Notify panic and then recover
		panic(nil)
	}

	fmt.Println("Pushed => ", objectToStk)
	stackElements[currentIndex] = objectToStk
	currentIndex++
}

// POP generic objects
func pop() (object any, CustomError any) {
	// Error Handling
	if currentIndex == -1 {
		return nil, errors.New("stack underflow occurred")
	}
	poppedElement := stackElements[currentIndex]
	currentIndex--
	return poppedElement, nil
}
