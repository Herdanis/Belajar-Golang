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
	address2 := &address1 // pointer ke address1
	address3 := &address1

	address2.City = "Bandung"

	// address2 = &Address{"Malang", "Jatim", "Indonesia"} // Mengubah semua isi Data mengunakan Operator &
	*address2 = Address{"Malang", "Jatim", "Indonesia"} // Semua mengacu pada Pointer ini Menggunakan Operator *

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	alamat1 := new(Address) // Membuat variabel baru dan kosong
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
