package main

import "fmt"

func main() {
	name := "qwerty"

	switch name {
	case "123":
		fmt.Println("angka")
	case "qwerty":
		fmt.Println("huruf")
	default:
		fmt.Println("unnone")
	}

	switch lang := len(name); lang > 2 {
	case true:
		fmt.Println("terlalu panjang")
	case false:
		fmt.Println("terlalu pendek")
	}
}
