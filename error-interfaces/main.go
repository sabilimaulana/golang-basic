package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

func Pembagian(nilai, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(9, 4)
	if err == nil {
		fmt.Println("Hasil =", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
