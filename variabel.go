package main

import "fmt"

func main() {
	var nama = "Eko Kurniawan" 
	var umur = 18
	country := "Indonesia"

	fmt.Println(nama, "Umur" , umur , country)

	nama = "Eko Kanedi"
	fmt.Println(nama,"Umur" , umur)

	var (
		firstname = "Eko"
		lastname = "kanedi"
	)
	
	fmt.Println(firstname , lastname)
}