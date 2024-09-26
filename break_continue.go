package main

import "fmt"

func main() {
	fmt.Println("==== Break dan Continue ====")

	//contoh break
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Println("Hallo Smua",i)
	}
	
	//contoh continue
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 { //semua angka gnap akan di lewati / continue
			continue
		}
		fmt.Println("Continue ganjil",i)
	}
	for i := 1; i <= 10; i++ {
		if i % 2 != 0 { //semua angka ganjil akan di lewati / continue
			continue
		}
		fmt.Println("Continue genap",i)
	}
}
