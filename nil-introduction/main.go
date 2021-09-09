package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	}

	return map[string]string{
		"name": name,
	}
}

func main() {
	data := newMap("Bili")
	// data := newMap("")

	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}
}
