package main

import "fmt"

type Address struct {
	kota, provinsi, negara string
}

func changeAddressToIndonesia(address *Address) {
	address.negara = "Indonesia"
}

func main() {
	address1 := Address{"Tegal", "Jawa Tengah", ""}

	fmt.Println(address1)
	changeAddressToIndonesia(&address1)
	fmt.Println(address1)
}
