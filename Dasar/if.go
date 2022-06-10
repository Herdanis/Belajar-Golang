package main

import "fmt"

func main() {
	name := "123"

	if name == "qwe" {
		fmt.Println("benar")
	} else if name == "123" {
		fmt.Println("lain")
	} else {
		fmt.Println("salah")
	}

	if leng := len(name); leng > 2 {
		fmt.Println("terlalu panjang")
	}
}
