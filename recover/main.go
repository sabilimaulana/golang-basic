package main

import "fmt"

//	wrong example

// func runApp(error bool) {
// 	fmt.Println("Menjalankan aplikasi")

// 	if error {
// 		panic("Error")
// 	}

// 	message := recover()
// 	fmt.Println("Terjadi error:", message)
// }

// This is the right one

func endApp() {
	fmt.Println("End...")

	message := recover()
	fmt.Println("Terjadi error:", message)
}

func runApp(error bool) {
	defer endApp()

	fmt.Println("Menjalankan aplikasi")

	if error {
		panic("ERROR 500")
	}
}

func main() {
	runApp(true)
}
