package main

import "fmt"

func fullName() (string, string, string) {
	return "Eko", "Kurniawan", "Kanedi"
}

func main() {
	var fistname, middlename, lastname = fullName()
	// var fistname, _, _ = fullName() // IGNORE DENGAN _
	fmt.Println(fistname)
	fmt.Println(middlename)
	fmt.Println(lastname)
}