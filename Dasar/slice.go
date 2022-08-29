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
		"okt",
		"nov",
		"des",
	}

	var slice1 = months[2:5]
	fmt.Println(slice1)

	months[3] = "diubah"
	fmt.Println(slice1)

	// slice menggunakan [ ] dan array menggunakan [leng]
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	slice2 := months[2:4]
	fmt.Println(slice2)

	// append slice
	slice3 := append(slice2, "TES")
	fmt.Println(slice3)
	slice3[1] = "bukan des"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// make new slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "qwe"
	newSlice[1] = "asd"

	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
	fmt.Println(len(newSlice))

	// copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

}
