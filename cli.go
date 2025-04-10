package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define command line flags for the two numbers
	num1 := flag.Float64("n1", 0.0, "First number to add")
	num2 := flag.Float64("n2", 0.0, "Second number to add")

	// Parse the command line arguments
	flag.Parse()

	// Calculate and print the result
	result := *num1 + *num2
	fmt.Printf("%.2f + %.2f = %.2f\n", *num1, *num2, result)
}