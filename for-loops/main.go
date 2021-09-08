package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Number", i)
	}

	// Another way
	counter := 1
	for counter <= 5 {
		fmt.Println("Perulangan ke", counter)
		counter++ // Don't forget to handle the counter
	}
}
