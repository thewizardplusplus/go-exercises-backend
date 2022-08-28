package main

import "fmt"

func productOfArray(numbers []int) int {
	product := 1
	for _, number := range numbers {
		product *= number
	}

	return product
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

	product := productOfArray(numbers)
	fmt.Print(product)
}
