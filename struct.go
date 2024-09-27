package main

import "fmt"

type CustomerPremium struct {
	Name, Addres string
	Age          int
}

func main() {
	var customer1 CustomerPremium
	customer1.Name = "Eko"
	customer1.Addres = "Indonesia"
	customer1.Age = 20
	fmt.Println(customer1)
}