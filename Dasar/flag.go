package main

import (
	"flag"
	"fmt"
)

func main() {
	// hasil dari flag adalah pointer
	host := flag.String("host", "localhost", "host name")
	user := flag.String("user", "root", "username")

	// perlu memanggil Parse() untuk menggunakan flag
	flag.Parse()

	// jika tidak menggunakan * hasil nya alamat memory

	fmt.Println("tidak menggunakan *")
	fmt.Println("host :", host)
	fmt.Println("host :", user)

	/* jika mengguna * hasil nya adalah value nya
	karena hasil dari flag adalah pointer */
	fmt.Println("\njika menggunakan *")
	fmt.Println("host :", *host)
	fmt.Println("host :", *user)

}
