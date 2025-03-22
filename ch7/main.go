package main

import (
	"bufio"
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
		fmt.Println("Choose an option:")
		fmt.Println("1. Perform a Calculation")
		fmt.Println("2. View Calculation History")
		fmt.Println("3. Exit")

		var option int
		fmt.Print("Enter choice (1-3): ")
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid choice.")
			continue
		}

		switch option {
		case 1:
			performCalculation(logger)
		case 2:
			viewHistory()
		case 3:
			fmt.Println("Thank you for using the calculator. Goodbye!")
			logger.Println("User exited the calculator.")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}

func performCalculation(logger *log.Logger) {
	var num1, num2 float64
	fmt.Print("Enter first number: ")
	_, err1 := fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	_, err2 := fmt.Scanln(&num2)

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid input. Please enter valid numbers.")
		return
	}

	fmt.Println("Choose an operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")

	var choice int
	fmt.Print("Enter choice (1-4): ")
	_, err3 := fmt.Scanln(&choice)
	if err3 != nil {
		fmt.Println("Invalid input. Please enter a valid choice.")
		return
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
	default:
		result = "Invalid choice. Please select a valid operation."
	}

	fmt.Println(result)
	logger.Println(result)
}

func viewHistory() {
	file, err := os.Open("calculator.log")
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	fmt.Println("Calculation History:")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading log file:", err)
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
