package main

import "fmt"

func person(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Nama": name,
		}
	}
}

func main() {
	person1 := person("Eko Kandesi")

	if person1 == nil {
		fmt.Println("Data Koosng")
	}else{
		fmt.Println(person1)
	}
}