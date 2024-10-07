package main

import "fmt"

type Adress struct {
	Kota, Provinsi, Negara string
}

func changeAddressToIndonesia(address1 *Adress) {
	address1.Negara = "Indonesia"
}

func main() {
	address1 := Adress{"Bandung", "Jawa Barat", ""}
	changeAddressToIndonesia(&address1)
	fmt.Println(address1)
}