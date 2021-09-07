package main

import "fmt"

func main() {
	// Menggunakan huruf kapital di awal kata seperti mendeklarasikan class pada javascript
	type NoKTP string

	var ktpBili NoKTP = "0812345678"

	fmt.Println(ktpBili)
	fmt.Println(NoKTP("345678"))
}
