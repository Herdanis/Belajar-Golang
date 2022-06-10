package main

import "fmt"

func main() {
	var months = [...]string{
		"jan",
		"feb",
		"mar",
		"apr",
		"mei",
		"jun",
		"jul",
		"agu",
		"sep",
	}

	var slice1 = months[2:5]
	fmt.Println(slice1)

	months[3] = "diubah"
	fmt.Println(slice1)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
