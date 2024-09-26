package main

import "fmt"

func sayHello(name string, filter func(string) string) {
	nameFilteres := filter(name)
	fmt.Println("Hallo saya", nameFilteres)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..
	}
	return name;
}

func main()  {
	sayHello("Anjing",spamFilter)
}