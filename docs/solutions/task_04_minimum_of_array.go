package main

import "fmt"

func findMinimumOfArray(numbers []int) int {
	minimum := numbers[0]
	for _, number := range numbers {
		if number < minimum {
			minimum = number
		}
	}

	return minimum
}

func main() {
	var numbers []int
	for {
		var number int
		if _, err := fmt.Scan(&number); err != nil {
			break
		}

		numbers = append(numbers, number)
	}

	minimum := findMinimumOfArray(numbers)
	fmt.Print(minimum)
}
