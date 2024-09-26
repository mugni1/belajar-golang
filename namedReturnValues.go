package main

import "fmt"

func getFullName2() (firstname string, middlename string, lastname string) {
	firstname = "eko"
	middlename = "kurniawan"
	lastname = "kanedi"
	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}