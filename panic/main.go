package main

import "fmt"

func endApp() {
	fmt.Println("End application")
}

func runApp(error bool) {
	defer endApp()
	fmt.Println("Menjalankan aplikasi")

	if error {
		panic("ERROR")
	}

	fmt.Println("Baris kode ini tidak akan berjalan")

}

func main() {
	runApp(true)
}
