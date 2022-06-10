package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "tes",
		"address": "qwe",
	}

	person["title"] = "Programmer"

	fmt.Println(person["title"])

	book := make(map[string]string)

	book["title"] = "asd"
	book["ups"] = "salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
