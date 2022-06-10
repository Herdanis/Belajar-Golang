package main

import "fmt"

func getName() (string, string) {
	return "123", "qwe"
}

func main() {
	a, b := getName()
	fmt.Println(a, b)

	_, x := getName()
	fmt.Println(x)
}
