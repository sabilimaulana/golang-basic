package main

import "fmt"

type Address struct {
	kota, provinsi, negara string
}

func main() {
	address1 := new(Address)

	address2 := address1

	address2.negara = "Indonesia"

	fmt.Println(address1, address2)
}
