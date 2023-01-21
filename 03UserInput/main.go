package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input lesson"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Kindly enter your address:")
	// coma ok // err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Address: ", input)
}
