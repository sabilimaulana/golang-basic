package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Sabili Maulana", "Sab"))
	fmt.Println(strings.Split("Sabili Maulana", " "))
	fmt.Println(strings.ToLower("Sabili Maulana"))
	fmt.Println(strings.ToUpper("Sabili Maulana"))
	fmt.Println(strings.Trim("    Sabili Maulana    ", " "))
	fmt.Println(strings.Compare("Sabili Maulana", "Sabili"))
	fmt.Println(strings.ReplaceAll("sabili maulana, sabili maulana", "sabili", "sbl"))
}
