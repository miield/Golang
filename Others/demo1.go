package main

import "fmt"

func main() {
	var a int
	a = 34
	fmt.Println("a = ", a)

	nameA := "Oyindamola"
	nameB := "oyindamola"
	nameC := "Oyindamola"

	result := nameA == nameB
	result1 := nameA == nameC

	fmt.Println(result)
	fmt.Println(result1)

}
