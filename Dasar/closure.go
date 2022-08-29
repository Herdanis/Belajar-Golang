package main

import "fmt"

func main() {
	counter := 0

	// in this func we can access variable counter
	increment := func() {
		fmt.Println("Increment")
		counter++

		/* but variable a cant access in main func
		because the variable a in this scope */
		a := "tes"
		fmt.Println(a)

		/* to use same name variable
		we can declare new variable example */
		counter := 0
		fmt.Println(counter)
	}

	increment()
	increment()

	fmt.Println(counter)
}
