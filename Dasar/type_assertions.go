package main

import (
	"fmt"
)

func random() interface{} {
	return 1
}
func main() {
	result := random()
	// newResult := result.(string)
	// fmt.Println(newResult)
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is Integer")
	default:
		fmt.Println("Unknown type")
	}
}
