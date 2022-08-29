package main

import "fmt"

//4. interface has func
type HasName interface {
	GetName() string
}

//5. inteface has call
func sayhello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

//3. struct func to use interface
func (person Person) GetName() string {
	return person.Name
}

func main() {
	//1. create variable struct
	var tes Person

	//2. input value "Tes" to struct
	tes.Name = "Tes"

	sayhello(tes)
}
