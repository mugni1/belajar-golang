package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	adress1 := Address{"GedeBage", "Bandung", "Indonesia"}
	adress2 := adress1 //copy dari address 1

	adress2.City = "Rancabolang"

	fmt.Println(adress1)
	fmt.Println(adress2)
}