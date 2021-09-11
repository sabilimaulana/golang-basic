package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())

	}

	float, err := strconv.ParseFloat("3.14", 16)
	if err == nil {
		fmt.Println(float)
	} else {
		fmt.Println(float)
	}

	booleanString := strconv.FormatBool(true)
	fmt.Println(strings.ToUpper(booleanString))
}
