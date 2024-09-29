package main

import "fmt"

type CustomerPremium struct {
	Name, Addres string
	Age          int
}

//contho2 
type Ktp struct{
	Nama, JenisKelamin, Ttl string
	Nik int
}

func main() {
	var customer1 CustomerPremium
	customer1.Name = "Eko"
	customer1.Addres = "Indonesia"
	customer1.Age = 20
	fmt.Println(customer1)

	var ktpAsep Ktp
	ktpAsep.Nik = 9123812921023
	ktpAsep.Nama = "Asep"
	ktpAsep.Ttl = "Bandung, 22-05-2007"

	fmt.Println(ktpAsep)
}