package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married(name string) { //gunakan * agar data asli ikut berubah
	man.Name = name + ". " + man.Name
}

func main() {
	eko := Man{"Eko"}
	eko.Married("Mr")

	nina := Man{"Nina"}
	nina.Married("Ms")

	fmt.Println(nina)
	fmt.Println(eko)
}