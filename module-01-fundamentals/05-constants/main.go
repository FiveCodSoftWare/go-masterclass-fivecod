package main

import "fmt"

func main() {
	// Constants cannot be changed
	const PiValue = 3.14159
	const AppName = "MyApp"
	const Version = "1.0"

	// Print constants
	fmt.Println("PiValue:", PiValue)
	fmt.Printf("App: %s v%s", AppName, Version)
}



