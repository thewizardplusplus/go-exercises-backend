package main

import "fmt"

func removeDuplicatesInPlace(numbers []int) int {
	insertIndex := 0
	for _, number := range numbers {
		var wasCopied bool
		for _, copyOfNumber := range numbers[:insertIndex] {
			if copyOfNumber == number {
				wasCopied = true
				break
			}
		}
		if wasCopied {
			continue
		}

		numbers[insertIndex] = number
		insertIndex++
	}

	return insertIndex
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

	cleanedNumberCount := removeDuplicatesInPlace(numbers)
	cleanedNumbers := numbers[:cleanedNumberCount]

	var cleanedNumbersForOutput []interface{}
	for _, number := range cleanedNumbers {
		cleanedNumbersForOutput = append(cleanedNumbersForOutput, number)
	}

	fmt.Print(cleanedNumbersForOutput...)
}
