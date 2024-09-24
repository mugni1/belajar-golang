package main

import "fmt"

func main() {
	//array
	//CONTOH 1
	var dataNama [3]string
	dataNama[0] = "asep"
	dataNama[1] = "ucep"
	dataNama[2] = "usuf"
	fmt.Println(dataNama[0])

	//CONTOH 2
	var values = [5]int{20,23,12,1,3,}
	fmt.Println("isi array =",values)
	fmt.Println("isi index ke 1 =",values[1])
	fmt.Println("panjang array ini =" ,len(values)) // len
	values[0] = 12; //edit or add
	fmt.Println("index ke 0 di ubah menjadi =" ,values[0])

	//NOTE : function array
	// 1. len(array) | untuk mendapatkan panjang array
	// 2. array[index] | untuk mendapat value dari index tsb
	// 3. array[index] = value | untuk menambah atau mengubah (jika sudah ada) value dari index tsb
}