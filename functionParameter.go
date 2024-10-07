package main

import "fmt"

func sayHello(nama string, umur int, alamat string) {
	fmt.Println("Hallo",nama,", Umur saya", umur, "Alamat " , alamat)
}

func main() {
	sayHello("Asep", 20, "Bandung");
}