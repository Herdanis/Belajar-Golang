package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke ", counter)
		counter++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("perulangan I ke ", i)
	}

	name := []string{"qwe", "asd", "zxc"}

	for i, value := range name {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "test"
	person["title"] = "programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
