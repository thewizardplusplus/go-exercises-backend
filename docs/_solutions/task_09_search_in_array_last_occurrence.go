package main

import "fmt"

func findLastNumberInArray(numbers []int, requiredNumber int) int {
	requiredIndex := -1
	for index, number := range numbers {
		if number == requiredNumber {
			requiredIndex = index
		}
	}

	return requiredIndex
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
	indexOfRequiredNumber := findLastNumberInArray(numbers, requiredNumber)
	fmt.Print(indexOfRequiredNumber)
}
