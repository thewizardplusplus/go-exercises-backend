package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	sum := add(x, y)
	fmt.Println(sum)
}
