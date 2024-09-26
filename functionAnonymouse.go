package main

import "fmt"

//function as parameter
func registerUser(name string, blacklist func(string) bool) { //arti blacklist function(parameter string) type nya bool
	if blacklist(name) == true {
		fmt.Println(name + " you are blocked")
	}else{
		fmt.Println("Welcome", name)
	}
}

func main() {
	//anonymouse function
	//contoh 1
	cekUser := func (name string) bool  {
		return name != "admin"
	}
	registerUser("Asep", cekUser)

	//contoh2
	registerUser("Anjay", func (namaUser string) bool  {
		return namaUser != "admin"
	})
}