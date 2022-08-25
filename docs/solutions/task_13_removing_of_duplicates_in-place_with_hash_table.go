package main

import "fmt"

func removeDuplicatesInPlace(numbers []int) int {
	insertIndex := 0
	index := make(map[int]struct{})
	for _, number := range numbers {
		if _, ok := index[number]; ok {
			continue
		}

		numbers[insertIndex] = number
		insertIndex++

		index[number] = struct{}{}
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
