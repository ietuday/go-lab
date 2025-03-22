package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s! Welcome to your first Go project.\n", name)

	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Println("Choose an operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	var choice int
	fmt.Print("Enter choice (1-4): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Printf("The sum of %d and %d is %d\n", num1, num2, addNumbers(num1, num2))
	case 2:
		fmt.Printf("The difference of %d and %d is %d\n", num1, num2, subtractNumbers(num1, num2))
	case 3:
		fmt.Printf("The product of %d and %d is %d\n", num1, num2, multiplyNumbers(num1, num2))
	case 4:
		if num2 != 0 {
			fmt.Printf("The division of %d by %d is %.2f\n", num1, num2, divideNumbers(num1, num2))
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Invalid choice. Please select a valid operation.")
	}
}

func addNumbers(a int, b int) int {
	return a + b
}

func subtractNumbers(a int, b int) int {
	return a - b
}

func multiplyNumbers(a int, b int) int {
	return a * b
}

func divideNumbers(a int, b int) float64 {
	return float64(a) / float64(b)
}
