package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Sabili")
	data.PushBack(17)
	data.PushBack("Maulana")

	fmt.Println(data)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
