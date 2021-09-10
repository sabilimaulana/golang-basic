package main

import "fmt"

type Address struct {
	kota, provinsi, negara string
}

func main() {
	address1 := Address{"Tegal", "Jawa Tengah", "Indonesia"}

	address2 := &address1

	address2.kota = "Pemalang"

	*address2 = Address{"Pati", "Jawa Tengah", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}
