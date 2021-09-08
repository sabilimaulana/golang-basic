package main

import "fmt"

func getBookData() (title string, author string, price int) {
	title = "Hidup Begitu Indah dan Hanya Itu yang Kita Punya"
	author = "Dea Anugrah"
	price = 50000

	return title, author, price // Urutan itu penting!
}

func main() {
	title, author, price := getBookData()

	fmt.Println(title, author, price)
}
