package main

import "fmt"

func logging() {
	fmt.Println("selesai")
}

func runApp(value int) {
	/* this mean the func must be call after execute done or error
	defer use in top line */
	defer logging()

	fmt.Println("run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApp(10)
}
