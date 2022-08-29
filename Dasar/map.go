package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"name":    "tes",
		"address": "qwe",
	}

	fmt.Println(person)

	//add new key and value
	person["title"] = "Programmer"
	fmt.Println(person)

	//make new map
	book := make(map[string]string)

	//change the value
	book["title"] = "asd"
	book["ups"] = "salah"
	fmt.Println(book)

	//delete map by key
	delete(book, "ups")
	fmt.Println(book)
}
