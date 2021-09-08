package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(4, 254, 712, 1, 54, 61, 42, 473, 131)

	fmt.Println("Total =", total)

	// Using slice as a value on variadic function
	numbers := []int{13, 54, 62, 32, 53, 68, 91}
	fmt.Println("Sum slice numbers =", sumAll(numbers...))
}
