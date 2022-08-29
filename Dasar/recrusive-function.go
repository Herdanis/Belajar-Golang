package main

import "fmt"

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}

	return result
}

// func yg memanggil dirinya sendiri
func factorialRecrusive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecrusive(value-1)
	}
}
func main() {

	loop := factorialLoop(5)
	fmt.Println(loop)

	loopR := factorialRecrusive(5)
	fmt.Println(loopR)
}
