package main

import "fmt"

type Blacklist func(name string) bool

func registerUser(name string, blacklist Blacklist) string {
	if blacklist(name) {
		return "You're has been blocked, " + name
	}

	return "Welcome " + name
}

func main() {
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	fmt.Println(registerUser("Sabili", blacklist))
	// Anonymous function
	fmt.Println(registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	}))
}
