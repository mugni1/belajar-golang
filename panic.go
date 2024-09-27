package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(eror bool)  {
	defer endApp()
	if eror == true {
		panic("errrrrooooorr")
	}
	fmt.Println("Aplikasi berjalan")
}

func main()  {
	runApp(false)
}