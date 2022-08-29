package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	// _ mean ignore the index
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	// variadic func use array/slice

	// example if slice not avalable
	total := sumAll(1, 3, 35, 6, 2)
	fmt.Println(total)

	// example if to use avalable slice
	slice := []int{10, 23, 13, 24, 25}
	//add new data
	slice = append(slice, 10)
	fmt.Println(slice)

	// to use slice in variadic use ... to convert slice to int
	total = sumAll(slice...)
	fmt.Println(total)

}
