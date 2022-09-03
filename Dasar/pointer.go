package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) { //parameter harus pointer ke Address
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Subang", "Jabar", "Indonesia"}
	// menggunakan '&' untuk mengarahkan value variable ke variable pertama
	address2 := &address1 // pointer ke address1
	address3 := &address1

	address2.City = "Bandung" // Hanya value dari city yg berubah

	/* Mengubah semua isi Data mengunakan Operator & tetapi membuat memory baru
	tidak mengacu ke value asli nya */
	address2 = &Address{"Malang", "Jatim", "Indonesia"}

	// Mengubah semua data dan semua yg mengacu pada address2 juga berubah ke data baru
	*address2 = Address{"Malang", "Jatim", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	alamat1 := new(Address) // Membuat variabel baru dan tidak memiliki data
	alamat1.City = "Jakarta"
	fmt.Println(alamat1)

	alamat := Address{
		City:     "Subang",
		Province: "Jabar",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat) // & mengubah data asli menjadi Pointer
	fmt.Println(alamat)
}
