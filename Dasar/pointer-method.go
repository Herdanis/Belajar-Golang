package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	// fmt.Println("Di Method", man.Name)
}

func main() {
	tes := Man{"Tes"}
	tes.Married()

	fmt.Println(tes.Name)
}
