package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// this name is struct methode or struct func
func (customer Customer) sayHello1(name string) {
	fmt.Println("hello", name, "is", customer.Name)
}

// without struct func
func sayHello1(customer Customer, name string) {
	fmt.Println("hello", name, "is", customer.Name)
}

func main() {

	// Make variable from struct
	var cus Customer
	cus.Name = "Tes"
	cus.Address = "qwe"
	cus.Age = 123

	// call value struct from variable
	fmt.Println(cus.Age)

	// this is how to call struct func
	cus.sayHello1("joko")

	// this is how to call without struct
	sayHello1(cus, "joko")

	/* 2 way to input data to struct
	Methode 1 */
	joko := Customer{
		Name:    "joko",
		Address: "qer",
		Age:     43,
	}

	fmt.Println(joko.Age)

	/* Methode 2 */
	budi := Customer{"budi", "123", 123}
	fmt.Println(budi)

}
