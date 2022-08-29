package main

import "fmt"

/* empty inteface to use return any
whict mean you can return any think
not like before must be declare the data type
is you can return any data type */
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	data := Ups(3)
	fmt.Println(data)

}
