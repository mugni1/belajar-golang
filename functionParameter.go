package main

import "fmt"

func sayHello(nama string, umur int) {
	fmt.Println("Hallo",nama,", Umur saya", umur)
}

func main() {
	sayHello("Asep", 20);
}