package main

import "fmt"

func main() {
	// Closures adalah kemampuan sebuah function untuk menggunakan data-data / variable-variable di luar scope function tersebut
	// Pada contoh kali ini adalah menggunakan variable counter

	counter := 0

	increment := func() {
		counter++
	}

	fmt.Println(counter)
	increment()
	increment()
	fmt.Println(counter)
}
