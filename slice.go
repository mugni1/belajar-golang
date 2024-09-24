package main

import "fmt"

func main() {
	var month = [...]string{
		"Jan","feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Aug", "Sept", "Okt", "Nov", "Des",
	}

	//NOTE | FUNCTION SLICE
	// 1. len(array/slice) | (length)
	// 2. cap(array/slice) | (capacity)
	// 3. append(slice, value) | mebuat slice baru dengan menambh data ke posisi terakhir slice jika kapasitas sudah penuh, maka akan membuat array baru
	// 4  make()
	// 5. copy(toSlice,fromSlice)
	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println("length",len(slice1)) // untuk mengetahui panjang | length slice1
	fmt.Println("Capacity",cap(slice1)) // untuk melihat total capasitas

	var slice2 = month[9:]
	var slice3 = append(slice2, "NewBulan"); //di sini dia membuat array baru jadi slice2 dan month tidak terubah sama sekalii
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(month)
	
	//mebuat new slice 
	var newSlice = make([]string, 2, 5) //make(array.TypeArray,length,capacity)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"
	fmt.Println("DATA NEW SLICE",newSlice)
	
	//copy data slice ke slice yang baru
	var copySlice = make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice) // copy(toSlice, fromSlice)
	fmt.Println("COPY FROM NEW SLICE",copySlice)
}