package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("U ar Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "123"
	}

	registerUser("123", blacklist)
	registerUser("qwe", blacklist)
}
