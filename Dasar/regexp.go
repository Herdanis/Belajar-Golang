package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("t([a-z])s")

	fmt.Println(regex.MatchString("tes"))
	fmt.Println(regex.MatchString("t1s"))
}
