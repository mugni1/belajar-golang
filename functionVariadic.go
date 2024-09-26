package main

import "fmt"

func getNameAndValue(number ...int) int {
	total := 0;
	for _, v := range number {
		total += v
	}
	return total
}

func main() {
	fmt.Println(getNameAndValue(10,20,30,40))
}