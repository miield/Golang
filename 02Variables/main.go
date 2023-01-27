package main

import "fmt"

// noVarKeyword := "hello there"// not allowed outside the method
var noVarKeyword2 int = 90000 //this is allowed

// ContractName First-letter-Capital variable are public variable. Constant cannot be changed
const ContractName string = "Golang" //line

// Variables Types

func main() {
	var age int = 30
	fmt.Println(age)
	fmt.Printf("Variable is a type: %T \n", age) // hello there

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

	// Default values and aliases
	var newVariable int // Declared not initialized + non garbage value
	fmt.Println(newVariable)
	fmt.Printf("Variable is a type: %T \n", newVariable)

	// No explicit type
	var noVariableType = false
	fmt.Println(noVariableType)
	// noVariableType = "Oyindamola" // cannot be reassign to a different type, lexer only accepts the variable type-value assigned
	fmt.Printf("Variable is a type: %T \n", noVariableType) // bool

	// No var keyword and type
	noVarKeyword := "hello there"
	fmt.Println(noVarKeyword)
	fmt.Printf("Variable is a type: %T \n", noVarKeyword)

	fmt.Println(ContractName)
	fmt.Printf("Variable is a type: %T \n", ContractName)

}

func main2() {
	var name string = "Oyindamola Abiola"
	fmt.Println(name)
	fmt.Printf("Variable is a type: %T \n", name)
}
