package main

import "fmt"

func main() {
	const a string = "tes"
	const b = "TES"
	const c = 1000

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	const (
		x = "test"
	)

	fmt.Println(x)
}
