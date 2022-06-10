package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	if message != nil {
		fmt.Println("Error dengan", message)
	}
	fmt.Println("Selesai")
}

func runApp1(error bool) {
	defer endApp()
	if error {
		panic("error")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp1(false)
}
