package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	/* recover func will catch the error info
	this func can work on defer func */
	message := recover()
	if message != nil {
		fmt.Println("Error dengan", message)
	}
	fmt.Println("Selesai")
}

func runApp1(error bool) {
	defer endApp()
	if error {
		// this mean if found error aplication must be stop
		panic("error")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp1(false)
}
