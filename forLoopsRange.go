package main

import "fmt"

func main() {
	//CONTOH PENGGUNAAN
	slice := []string{"Eko", "Agus", "Iki", "tes"}

	// FOR SLICE
	fmt.Println("FOR RANGE UNTUK SLICE")
	//menggunakan for range | untuk slice dan array
	for _, value := range slice { //jika tidak ingin ada indexnya gunakan _
		// fmt.Println("index", index, "=",  value)
		fmt.Println(value)
	}

	//FOR MAP
	map1 := make(map[string]string)
	map1["title"] = "Ini title"
	map1["isi"] = "lorem ipsum dolor sit amet"
	fmt.Println("FOR RANGE UNTUK MAP")
	for _, v := range map1 {
		fmt.Println(v)
	}

	//FOR ARRAU
	arr1 := [...]string{"eko", "kurniawan", "kanedi"}
	fmt.Println("FOR RANGE UNTUK ARRAY")
	for _, v := range arr1 {
		fmt.Println(v)
	}
}