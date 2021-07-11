package main

import "fmt"

func main() {
    /* Array for sorting */
	var numbers = []int{5, 6, 8, 10, 23, 6, 7, 7, 7}
	var i, j int

	for i = 0; i < len(numbers)-1; i++ {
		for j = i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				var temp = numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}

	for j = 0; j < len(numbers); j++ {
		fmt.Printf("%d,", numbers[j])
	}
}
