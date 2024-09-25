package main

import "fmt"

func main() {
	nama := "Ucup"
	switch nama {
	case "Eko":
		fmt.Println("Nama Anda adalah Eko");
		fmt.Println("Nama Anda adalah Eko");
	case "Ucup":
		fmt.Println("Nama Anda adalah ucup")
	default:
		fmt.Println("Nama yang anda masukan tidak di kenal")
	}
		
	//sort switch statement
	switch length := len(nama); {
	case length > 5:
		fmt.Println("Panjang",nama,"lebih dari lima")
	case length < 5:
	    fmt.Println("panjang",nama,"Kurang dari lima")
	default :
	    fmt.Println("Tidak benar")
    }

	//switch tanpa kondisi
	var length2 = len(nama)
	switch  {
	case length2 > 5:
		fmt.Println("lebih dari 5")
	case length2 < 5:
		fmt.Println("kurang dari 5")
	default:
		fmt.Println("Asu")
	}
}