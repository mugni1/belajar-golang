package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	adress1 := Address{"GedeBage", "Bandung", "Indonesia"}
	adress2 := adress1 //copy dari address 1
	adress3 := &adress1 //references dari address 1 | & di sini untuk mereference ke addrees 1

	adress2.City = "Rancabolang"
	adress3.City = "Rancananggung"

	*adress3 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println("Adress 1", adress1)
	fmt.Println("Adress 2",adress2)
	fmt.Println("Adress 3",adress3)
	
	// new 
	var adress4 *Address = new(Address)
	adress4.City = "Jakarta"
	fmt.Println(adress4)
}