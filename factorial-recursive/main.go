package main

import "fmt"

func factorialRecursive(number int) int {
	if number == 1 {
		return number
	}

	return number * factorialRecursive(number-1)
}

func main() {
	fmt.Println(factorialRecursive(5))
}
