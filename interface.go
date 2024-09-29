package main

import "fmt"

type HasName interface {
	GetName() string //
}
func sayHellooo(hasname HasName) {
	fmt.Println("Hallo ini adalah", hasname.GetName())
}

//implementasi (1) dengan struct
type Person struct{
	Nama string
}
func (person Person) GetName() string { //
	return person.Nama
}

// implementasi (2)
type Animall struct{
	Nama string
}
func (animal Animall) GetName() string  {
	return  animal.Nama
}

func main() {
	var eko Person
	eko.Nama = "Eko"
	sayHellooo(eko)

	anjing := Animall{
		Nama: "Bobi",
	}
	sayHellooo(anjing)
}