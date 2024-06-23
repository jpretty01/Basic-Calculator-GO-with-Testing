package main

import (
	"fmt"
)

// Add function to add two numbers together
func add(x int, y int) int {
	return x + y
}

// Subtract function to subtract one number from another
func subtract(x int, y int) int {
	return x - y
}

// Multiply function to multiply two numbers together
func multiply(x int, y int) int {
	return x * y
}

// Divide function to divide one number by another (handling division by zero)
func divide(x int, y int) float64 {
	if y == 0 {
		fmt.Println("Error: Division by zero is not allowed.")
		return 0
	}
	return float64(x) / float64(y)
}

func main() {
	fmt.Println("Simple Calculator")
	fmt.Println("------------------")

	for {
		var x, y int
		fmt.Print("Enter first number: ")
		fmt.Scan(&x)
		fmt.Print("Enter second number: ")
		fmt.Scan(&y)

		fmt.Println("Addition:", add(x, y))
		fmt.Println("Subtraction:", subtract(x, y))
		fmt.Println("Multiplication:", multiply(x, y))
		fmt.Println("Division:", divide(x, y))

		var answer string
		fmt.Print("Continue? [y/n]: ")
		fmt.Scan(&answer)

		if answer != "y" && answer != "Y" {
			break
		}
	}
}
