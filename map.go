package main

import (
	"fmt"
)

func main() {
	//NOTE : FUNCTION MAP
	// 1. len(map) mengetahui length
	// 2. map[key] mengambil value
	// 3. map[key] = value | mengubah jika sudah ada value / menambah jika belum ada value
	// 4. make(map[typekey]typevalue) | membuat map baru
	// 5. delete(map, key) | menghapus sebagian key dari map

	map1 := map[string]string{
		"name":  "Asep",
		"kelas": "XII RPL II",
	}
	//add
	map1["title"] = "programmer"

	fmt.Println("Length=",len(map1))
	fmt.Println("Value Key Name =",map1["name"])
	delete(map1, "kelas"); //delete key "kelas"
	fmt.Println("Key And Value =", map1)

	//membuat map baru
	book := make(map[string]string)
	book["judul"] = "The Rock"
	book["isi"] = "lorem ipsum dolor sit amet"
	fmt.Println(book)
}