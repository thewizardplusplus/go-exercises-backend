package main

import "fmt"

func findIndexOfMinimumOfArray(numbers []int) int {
	minimum := numbers[0]
	indexOfMinimum := 0
	for index, number := range numbers {
		if number < minimum {
			minimum = number
			indexOfMinimum = index
		}
	}

	return indexOfMinimum
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

	indexOfMinimum := findIndexOfMinimumOfArray(numbers)
	fmt.Print(indexOfMinimum)
}
