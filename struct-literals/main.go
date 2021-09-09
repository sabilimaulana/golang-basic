package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func main() {
	sbl := Customer{
		name:    "Sabili",
		address: "Tegal",
		age:     18,
	}

	fmt.Println("Hello! my name is", sbl.name)
	fmt.Println(sbl.age, "years old!")
	fmt.Println("I am living in", sbl.address)

	zeus := Customer{"Zeus", "Olympics", 100}
	fmt.Println(zeus)
}
