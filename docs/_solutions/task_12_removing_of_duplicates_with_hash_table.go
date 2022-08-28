package main

import "fmt"

func removeDuplicates(numbers []int) []int {
	var copyOfNumbers []int
	index := make(map[int]struct{})
	for _, number := range numbers {
		if _, ok := index[number]; ok {
			continue
		}

		copyOfNumbers = append(copyOfNumbers, number)
		index[number] = struct{}{}
	}

	return copyOfNumbers
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

	cleanedNumbers := removeDuplicates(numbers)

	var cleanedNumbersForOutput []interface{}
	for _, number := range cleanedNumbers {
		cleanedNumbersForOutput = append(cleanedNumbersForOutput, number)
	}

	fmt.Print(cleanedNumbersForOutput...)
}
