package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	data, err := Pembagian(1, 0)
	if err == nil {
		fmt.Println("hasil", data)
	} else {
		fmt.Println("error", err)
	}
}
