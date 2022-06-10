package main

import "fmt"

func main() {
	type noKtp string

	var noKtpTest noKtp = "12312413451"
	fmt.Println(noKtpTest)
}
