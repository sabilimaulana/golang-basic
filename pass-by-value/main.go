package main

import "fmt"

type Address struct {
	kota, provinsi, negara string
}

func main() {
	address1 := Address{"Tegal", "Jawa Tengah", "Indonesia"}

	address2 := address1

	address2.kota = "Pati"

	fmt.Println(address1) // address 1 tidak berubah
	fmt.Println(address2)

}
