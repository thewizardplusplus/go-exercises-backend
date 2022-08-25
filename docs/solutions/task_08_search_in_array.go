package main

import "fmt"

func findNumberInArray(numbers []int, requiredNumber int) int {
	for index, number := range numbers {
		if number == requiredNumber {
			return index
		}
	}

	return -1
}

func main() {
	var totalNumbers []int
	for {
		var number int
		if _, err := fmt.Scan(&number); err != nil {
			break
		}

		totalNumbers = append(totalNumbers, number)
	}

	numbers, requiredNumber :=
		totalNumbers[:len(totalNumbers)-1], totalNumbers[len(totalNumbers)-1]
	indexOfRequiredNumber := findNumberInArray(numbers, requiredNumber)
	fmt.Print(indexOfRequiredNumber)
}
