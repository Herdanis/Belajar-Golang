package main

import "fmt"

func getName() (string, string) {
	return "123", "qwe"
}

func main() {
	a, b := getName()
	fmt.Println(a, b)

	// to ignore return value use _
	_, x := getName() // this mean return first return value not use
	fmt.Println(x)
}
