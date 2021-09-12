package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`s([a-z])l`)

	fmt.Println(regex.MatchString("sbl"))
	fmt.Println(regex.MatchString("mvl"))
	fmt.Println(regex.MatchString("syl"))
	fmt.Println(regex.FindAllString("shl ssl sql sml jol", 10))

}
