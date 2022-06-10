package main

import "fmt"

func getGoodBye(name string) string {
	return "good bye " + name
}

func main() {
	sayGoodbye := getGoodBye

	result := sayGoodbye("qwe")
	fmt.Println(result)
}
