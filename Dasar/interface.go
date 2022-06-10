package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayhello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var tes Person
	tes.Name = "Tes"
	sayhello(tes)
}
