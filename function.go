package main

import "fmt"

//function
func sayHello() {
	fmt.Println("Hallo Semuanya")
}

//main function
func main() {
	for i := 1; i <= 3; i++ {
		sayHello()
	}
}