package main

import "fmt"

func sayHelloFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "123" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("123", spamFilter)
	sayHelloFilter("qwe", spamFilter)
}
