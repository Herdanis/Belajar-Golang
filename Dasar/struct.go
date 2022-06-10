package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello1(name string) {
	fmt.Println("hello", name, "is", customer.Name)
}

func sayHello1(customer Customer, name string) {
	fmt.Println("hello", name, "is", customer.Name)
}

func main() {
	var cus Customer
	cus.Name = "Tes"
	cus.Address = "qwe"
	cus.Age = 123

	fmt.Println(cus.Age)

	cus.sayHello1("joko")
	sayHello1(cus, "joko")

	joko := Customer{
		Name:    "joko",
		Address: "qer",
		Age:     43,
	}

	fmt.Println(joko.Age)

	budi := Customer{"budi", "123", 123}
	fmt.Println(budi)

}
