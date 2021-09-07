package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Different between array and slice

	array := [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	slice := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	fmt.Println(array, reflect.TypeOf(array), "<- This is an array!")
	fmt.Println(slice, reflect.TypeOf(slice), "<- This is a slice!")

	// Be careful
}
