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

	// arrays
	var simple_array = [3]int{23, 45, 98}
	inferred_length_array := [...]int{38, 52, 58, 9}
	var yennefer = [5]string{"Yennefer", "is", "a", "good", "girl"}
	fmt.Println(simple_array)
	fmt.Println(inferred_length_array)
	fmt.Println(yennefer)
}
