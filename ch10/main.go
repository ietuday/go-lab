package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	fmt.Println("Choose an operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Exponentiation")
	fmt.Println("6. Square Root")
	fmt.Println("7. Factorial")
	fmt.Println("8. Modulus (Remainder)")
	fmt.Println("9. Absolute Value")

	var choice int
	fmt.Print("Enter choice (1-9): ")
	_, err3 := fmt.Scanln(&choice)
	if err3 != nil {
		fmt.Println("Invalid input. Please enter a valid choice.")
		return
	}

	var result string
	switch choice {
	case 1:
		fmt.Print("Enter second number: ")
		_, err2 := fmt.Scanln(&num2)
		if err2 != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}
		res := addNumbers(num1, num2)
		result = fmt.Sprintf("The sum of %.2f and %.2f is %.2f", num1, num2, res)
	case 2:
		fmt.Print("Enter second number: ")
		_, err2 := fmt.Scanln(&num2)
		if err2 != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}
		res := subtractNumbers(num1, num2)
		result = fmt.Sprintf("The difference of %.2f and %.2f is %.2f", num1, num2, res)
	case 3:
		fmt.Print("Enter second number: ")
		_, err2 := fmt.Scanln(&num2)
		if err2 != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}
		res := multiplyNumbers(num1, num2)
		result = fmt.Sprintf("The product of %.2f and %.2f is %.2f", num1, num2, res)
	case 4:
		fmt.Print("Enter second number: ")
		_, err2 := fmt.Scanln(&num2)
		if err2 != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}
		if num2 != 0 {
			res := divideNumbers(num1, num2)
			result = fmt.Sprintf("The division of %.2f by %.2f is %.2f", num1, num2, res)
		} else {
			result = "Error: Division by zero is not allowed."
		}
	case 5:
		fmt.Print("Enter exponent: ")
		_, err2 := fmt.Scanln(&num2)
		if err2 != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}
		res := exponentiate(num1, num2)
		result = fmt.Sprintf("%.2f raised to the power of %.2f is %.2f", num1, num2, res)
	case 6:
		result = fmt.Sprintf("The square root of %.2f is %.2f", num1, squareRoot(num1))
	case 7:
		result = fmt.Sprintf("The factorial of %d is %d", int(num1), factorial(int(num1)))
	case 8:
		fmt.Print("Enter second number: ")
		_, err2 := fmt.Scanln(&num2)
		if err2 != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			return
		}
		result = fmt.Sprintf("The remainder of %.2f divided by %.2f is %.2f", num1, num2, modulus(num1, num2))
	case 9:
		result = fmt.Sprintf("The absolute value of %.2f is %.2f", num1, absoluteValue(num1))
	default:
		result = "Invalid choice. Please select a valid operation."
	}

	fmt.Println(result)
	logger.Println(result)
}

func modulus(a, b float64) float64         { return math.Mod(a, b) }
func absoluteValue(a float64) float64      { return math.Abs(a) }
func addNumbers(a, b float64) float64      { return a + b }
func subtractNumbers(a, b float64) float64 { return a - b }
func multiplyNumbers(a, b float64) float64 { return a * b }
func divideNumbers(a, b float64) float64   { return a / b }
func exponentiate(a, b float64) float64    { return math.Pow(a, b) }
func squareRoot(a float64) float64         { return math.Sqrt(a) }

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func viewHistory() {
	file, err := os.Open("calculator.log")
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Calculation History:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
