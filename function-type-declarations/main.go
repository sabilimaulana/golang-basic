package main

import "fmt"

// Function Type Declaration
type Filter func(name string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name
}

func main() {
	sayHelloWithFilter("Sabili", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

}
