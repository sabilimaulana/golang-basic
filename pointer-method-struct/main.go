package main

import "fmt"

type Man struct {
	name string
}

func (man *Man) Married() {
	man.name = "Mr " + man.name
}

func main() {
	bili := Man{"Sabili"}
	bili.Married()
	fmt.Println(bili)
}
