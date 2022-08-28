package main

import "fmt"

func findMaximumOfArray(numbers []int) int {
	maximum := numbers[0]
	for _, number := range numbers {
		if number > maximum {
			maximum = number
		}
	}

	return maximum
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

	maximum := findMaximumOfArray(numbers)
	fmt.Print(maximum)
}
