package main

import "fmt"

// Variables Types

func main() {
	var age int = 30
	fmt.Println(age)
	fmt.Printf("Variable is a type: %T \n", age)

	var name string = "Oyindamola Abiola"
	fmt.Println(name)
	fmt.Printf("Variable is a type: %T \n", name)

	var isFemale bool = true
	fmt.Println(isFemale)
	fmt.Printf("Variable is a type: %T \n", isFemale)

	var smallVal uint8 = 255 // 256 overflows
	fmt.Println(smallVal)
	fmt.Printf("Variable is a type: %T \n", smallVal)

	var smallFloat float32 = 255.5678268289398 // float32 5decimals : float64 long decimal
	fmt.Println(smallFloat)
	fmt.Printf("Variable is a type: %T \n", smallFloat)
}

// func main2() {
// 	var name string = "Oyindamola Abiola"
// 	fmt.Println(name)
// 	fmt.Printf("Variable is a type: %T \n", name)
// }
