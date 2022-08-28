package main

import "fmt"

func removeDuplicates(numbers []int) []int {
	var copyOfNumbers []int
	for _, number := range numbers {
		var wasCopied bool
		for _, copyOfNumber := range copyOfNumbers {
			if copyOfNumber == number {
				wasCopied = true
				break
			}
		}
		if wasCopied {
			continue
		}

		copyOfNumbers = append(copyOfNumbers, number)
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
