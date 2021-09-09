package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func main() {
	var sbl Customer

	sbl.name = "Sabili"
	sbl.address = "Tegal"
	sbl.age = 18

	fmt.Println(sbl.name)
	fmt.Println(sbl)
}
