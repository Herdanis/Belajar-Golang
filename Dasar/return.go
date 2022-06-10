package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "hello 00"
	} else {
		return "hello " + name
	}
}
func main() {
	fmt.Println(getHello("erba"))
}
