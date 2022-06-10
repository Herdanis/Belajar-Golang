package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "123"
	middleName = "qwe"
	lastName = "asd"
	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
