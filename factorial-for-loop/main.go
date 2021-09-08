package main

import "fmt"

func factorialLoop(number int) int {
	result := 1
	for i := number; i > 0; i-- {
		result *= i
	}

	return result
}

func main() {
	fmt.Println(factorialLoop(5))
}
