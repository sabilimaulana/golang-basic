package main

import "fmt"

func main() {
	var fruits = [4]string{
		"Apple",
		"Banana",
		"Watermelon",
		"Mango",
	}

	fmt.Println(fruits)
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])
	fmt.Println(fruits[2])
	fmt.Println(fruits[3])

	// Another array declaration
	var values [2]int
	values[0] = 17
	values[1] = 2003
	fmt.Println((values))

	// Another
	// If you don't know how your array length will
	colors := [...]string{"Red", "Yellow", "Blue", "White", "Black", "Lime"}
	fmt.Println(colors[len(colors)-1])
}
