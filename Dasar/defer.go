package main

import "fmt"

func logging() {
	fmt.Println("selesai")
}

func runApp(value int) {
	defer logging()
	fmt.Println("run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApp(10)
}
