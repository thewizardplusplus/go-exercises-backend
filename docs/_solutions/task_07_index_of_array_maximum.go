package main

import "fmt"

func findIndexOfMaximumOfArray(numbers []int) int {
	maximum := numbers[0]
	indexOfMaximum := 0
	for index, number := range numbers {
		if number > maximum {
			maximum = number
			indexOfMaximum = index
		}
	}

	return indexOfMaximum
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

	indexOfMaximum := findIndexOfMaximumOfArray(numbers)
	fmt.Print(indexOfMaximum)
}
