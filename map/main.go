package main

import "fmt"

func main() {
	identity := map[string]string{
		"name":      "Sabili Maulana",
		"hairColor": "black",
	}

	fmt.Println(identity)
	fmt.Println(identity["name"])
	fmt.Println(identity["hairColor"])

	identity["nationality"] = "Indonesia"
	fmt.Println(identity)

	// fmt.Println(identity.name) <- error

	book := make(map[string]string)

	book["title"] = "A Great Book!"
	book["author"] = "An Cool Author"
	book["wrong"] = "ups"
	fmt.Println(book)
	delete(book, "wrong")
	fmt.Println(book)

}
