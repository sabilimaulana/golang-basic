package main

import "fmt"

func main() {
	colors := []string{"Black", "White", "Lime", "Green"}

	for index, color := range colors {
		fmt.Println("Index ke", index, "=", color)
	}

	// Break VS Continue
	files := []string{"data a", "data b", "error", "data d"}

	for index, file := range files {
		if file == "error" {
			break
		}

		fmt.Println(index, file)
	}

	values := []int{90, 88, 77, 34, 78, 54, 87}

	for _, value := range values {
		if value < 70 {
			continue
		}

		fmt.Println("Value", value, "passed")
	}
}
