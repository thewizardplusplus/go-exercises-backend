package main

import "fmt"

func sumOfArray(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}

	return sum
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

	sum := sumOfArray(numbers)
	fmt.Print(sum)
}
