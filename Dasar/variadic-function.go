package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(1, 3, 35, 6, 2)
	fmt.Println(total)

	slice := []int{10, 23, 13, 24, 25}
	total = sumAll(slice...)
	fmt.Println(total)

}
