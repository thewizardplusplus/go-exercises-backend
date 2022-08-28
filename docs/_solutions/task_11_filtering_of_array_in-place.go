package main

import "fmt"

func filterArrayInPlace(numbers []int, unwantedNumber int) int {
	insertIndex := 0
	for _, number := range numbers {
		if number == unwantedNumber {
			continue
		}

		numbers[insertIndex] = number
		insertIndex++
	}

	return insertIndex
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

	numbers, unwantedNumber :=
		totalNumbers[:len(totalNumbers)-1], totalNumbers[len(totalNumbers)-1]
	filteredNumberCount := filterArrayInPlace(numbers, unwantedNumber)
	filteredNumbers := numbers[:filteredNumberCount]

	var filteredNumbersForOutput []interface{}
	for _, number := range filteredNumbers {
		filteredNumbersForOutput = append(filteredNumbersForOutput, number)
	}

	fmt.Print(filteredNumbersForOutput...)
}
