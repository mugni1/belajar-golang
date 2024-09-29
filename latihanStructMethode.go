package main

import "fmt"

type Animal struct {
	Nama, Warna string
	Panjang     int
}

// add methode for struct
func (animal Animal) bersuara(suara string) string {
	return animal.Nama + " Suaranya " + suara
}

func main() {
	var anjing Animal
	anjing.Nama = "Bone"
	anjing.Warna = "Putih"
	anjing.Panjang = 1
	fmt.Println(anjing.bersuara("Guk Guk"))
}