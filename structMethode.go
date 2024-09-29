package main

import "fmt"

//struct baru
type Customer struct {
	Nama, Kelas string
	Umur        int
}

//struct methode
func (customer Customer) sayHi(name string) {
	fmt.Println("Hallo, " + name + " Saya " + customer.Nama)
}

func main()  {
	//definiskan Cusstomer
	var eko Customer
	eko.Nama = "Eko"
	eko.sayHi("joko")
}