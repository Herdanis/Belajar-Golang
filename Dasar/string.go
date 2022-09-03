package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(len("Tes"))
	fmt.Println("Tes"[0])

	// return boolean
	fmt.Println(strings.Contains("tes qwe asd", "qwe"))

	// return slice
	fmt.Println(strings.Split("tes qwe asd", " "))

	// berubah menjadi upper/lower case
	fmt.Println(strings.ToUpper("tes qwe asd"))
	fmt.Println(strings.ToLower("tEs Qwe ASd"))

	// menghapus karakter di kiri dan kanan
	fmt.Println(strings.Trim("   tEs Qwe ASd ", " "))

}
