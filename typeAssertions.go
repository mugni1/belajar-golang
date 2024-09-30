package main

import "fmt"

func random() interface{} {
	return true
}

// KONVERTION
func main() {
	result := random()
	// resultString := result.(string) // konversi isi dari interface ke string
	// fmt.Println(resultString)

	// resultInt := result.(int) // panic | pastikan data yg di konversi sesuai
	// fmt.Println(resultInt)
	switch result.(type) {
	case string:
		fmt.Println("Result", result , "Is String")
	case int:
		fmt.Println("Result", result, "Is Int")
	case bool:
		fmt.Println("Result", result, "Is bool")
	default :
	    fmt.Println("Unknow Type")
	}
}