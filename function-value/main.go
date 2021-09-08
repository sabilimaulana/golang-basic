package main

import "fmt"

func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	fmt.Println(getGoodbye("sbl"))

	goodbye := getGoodbye

	fmt.Println(goodbye("Sabili"))

}
