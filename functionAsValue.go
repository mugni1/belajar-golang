package main

import "fmt"

func sayGodBy(nama string) string {
	return "GoodBye " + nama
}

func main() {
	godbye := sayGodBy
	fmt.Println(godbye("Asep"))
}