package main

import "fmt"

// example 'filter' is name of func parametes

/* can use type to declare the func for example :
type Filter func(string, string, int, int) string */
func sayHelloFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)
}

// func for filtering
func spamFilter(name string) string {
	if name == "123" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("123", spamFilter)
	filter := spamFilter
	sayHelloFilter("qwe", filter)
}
