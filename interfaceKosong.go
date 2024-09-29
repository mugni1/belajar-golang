package main

import "fmt"

func ups(param int) interface{} {
	if param == 1 {
		return 1
	} else if param == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	fmt.Println(ups(3))
}