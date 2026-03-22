package main

import "fmt"

func main() {
	// Explicit type declaration
	var text string = "Hello"
	var number int = 100
	var decimal float64 = 3.14
	var truth bool = true

	// Implicit type declaration
	var anotherText = "World"

	// Print with labels
	fmt.Println("String:", text)
	fmt.Println("Int:", number)
	fmt.Println("Float:", decimal)
	fmt.Println("Bool:", truth)
	fmt.Println("Another Text:", anotherText)
}
