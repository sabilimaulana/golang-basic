package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

// Struct method
func (customer Customer) SayHello() {
	fmt.Println("Hello", customer.name)
}

func main() {
	sbl := Customer{"Sabili", "Tegal", 18}

	sbl.SayHello()
}
