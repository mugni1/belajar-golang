package main

import "fmt"

func main() {
	//cara 1
	result := 1
	for result <= 3 {
		fmt.Println("Hallo Dunia" , result)
		result++;
	}
	
	//cara 2
	for i := 1; i <= 3; i++ {
		fmt.Println("Hallo World",i)
	}

	//penggunaab
	slice := []string{"Eko","Agus","Iki","tes"}
	
	// menggunakan for loops
	for i := 0; i < len(slice); i++ {
		fmt.Println(i+1, slice[i])
	}

	//menggunakan for range | untuk slice dan array
	for index, value := range slice { 
		fmt.Println("index", index, "=",  value)
	}
}