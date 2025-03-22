package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("calculator.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	defer file.Close()
	logger := log.New(file, "LOG: ", log.Ldate|log.Ltime)

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s! Welcome to your first Go project.\n", name)
	logger.Printf("User %s started the calculator.\n", name)

	for {
		var num1, num2 float64
		fmt.Print("Enter first number: ")
		_, err1 := fmt.Scanln(&num1)
		fmt.Print("Enter second number: ")
		_, err2 := fmt.Scanln(&num2)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid input. Please enter valid numbers.")
			continue
		}

		fmt.Println("Choose an operation:")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter choice (1-5): ")
		_, err3 := fmt.Scanln(&choice)

		if err3 != nil {
			fmt.Println("Invalid input. Please enter a valid choice.")
			continue
		}

		var result string
		switch choice {
		case 1:
			res := addNumbers(num1, num2)
			result = fmt.Sprintf("The sum of %.2f and %.2f is %.2f", num1, num2, res)
		case 2:
			res := subtractNumbers(num1, num2)
			result = fmt.Sprintf("The difference of %.2f and %.2f is %.2f", num1, num2, res)
		case 3:
			res := multiplyNumbers(num1, num2)
			result = fmt.Sprintf("The product of %.2f and %.2f is %.2f", num1, num2, res)
		case 4:
			if num2 != 0 {
				res := divideNumbers(num1, num2)
				result = fmt.Sprintf("The division of %.2f by %.2f is %.2f", num1, num2, res)
			} else {
				result = "Error: Division by zero is not allowed."
			}
		case 5:
			fmt.Println("Thank you for using the calculator. Goodbye!")
			logger.Println("User exited the calculator.")
			return
		default:
			result = "Invalid choice. Please select a valid operation."
		}

		fmt.Println(result)
		logger.Println(result)
	}
}

func addNumbers(a float64, b float64) float64 {
	return a + b
}

func subtractNumbers(a float64, b float64) float64 {
	return a - b
}

func multiplyNumbers(a float64, b float64) float64 {
	return a * b
}

func divideNumbers(a float64, b float64) float64 {
	return a / b
}
