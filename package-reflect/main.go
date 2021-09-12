package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

func main() {
	sample := Sample{"Sabili"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	fmt.Println(structField.Name)
}
