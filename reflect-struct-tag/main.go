package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func main() {
	sample := Sample{"Sabili"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	required := structField.Tag.Get("required")

	fmt.Println(required)
}
