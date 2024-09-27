package main

import "fmt"

func logging() {
	fmt.Println("Selesai di jalankan")
}

func runApplication()  {
	defer logging()
	fmt.Println("Aplikasi di jalankan")
}

func main() {
	runApplication()
}