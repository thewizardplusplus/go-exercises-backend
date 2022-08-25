package main

import "fmt"

func filterArray(numbers []int, unwantedNumber int) []int {
	var filteredNumbers []int
	for _, number := range numbers {
		if number != unwantedNumber {
			filteredNumbers = append(filteredNumbers, number)
		}
	}

	return filteredNumbers
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
	filteredNumbers := filterArray(numbers, unwantedNumber)

	var filteredNumbersForOutput []interface{}
	for _, number := range filteredNumbers {
		filteredNumbersForOutput = append(filteredNumbersForOutput, number)
	}

	fmt.Print(filteredNumbersForOutput...)
}
