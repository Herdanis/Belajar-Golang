package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "tes"
	names[1] = "asd"
	names[2] = "qwe"

	fmt.Println(names[2])

	var values = [3]int{
		49,
		27,
		50,
	}

	fmt.Println(values[0])

	fmt.Println(len(names))
}
