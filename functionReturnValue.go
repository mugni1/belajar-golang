package main

import "fmt"

func sayHallo(nama string) string { //string yang di luar adalah tipe data return
	return "Hallo saya "+ nama
}

func cekUsername(username string) string {
	if username == "asep22" {
		return "Username benar"
	}

	return "Username tidak ada"
}

func main() {
	fmt.Println(sayHallo("Asep"))
	fmt.Println(cekUsername("asep22"))
}