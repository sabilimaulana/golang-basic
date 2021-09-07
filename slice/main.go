package main

import "fmt"

func main() {
	nations := [...]string{"Indonesia", "Malaysia", "USA", "UK", "German", "Japan", "Spain"}
	slice := nations[:]

	fmt.Println(slice)
	fmt.Println(slice[2])

	fmt.Println("Capacity", cap(slice)) // Capacity based on array nations
	fmt.Println("Length", len(slice))

	slice = append(slice, "Russia", "Italy")
	fmt.Println(slice)

	fmt.Println("Capacity", cap(slice)) // Capacity based on array nations
	fmt.Println("Length", len(slice))

	// Make slice -> make([]typedata, length, capacity)
	slice2 := make([]string, 3, 5)
	slice2[0] = "Komodo"
	slice2[1] = "Cat"
	slice2[2] = "Butterfly"
	// slice2[3] = "Buffalo" error

	fmt.Println(slice2)
}
