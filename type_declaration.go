package main

import "fmt"

func main() {
	//alias untuk type data
	type NoKTP string

	var ktpEko NoKTP = "12933223"
	fmt.Println(ktpEko)
}