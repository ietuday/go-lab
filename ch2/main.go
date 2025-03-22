package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s! Welcome to your first Go program.\n", name)

	var num1, num2 int
	fmt.Print("Enter two numbers: ")
	fmt.Scanln(&num1, &num2)
	fmt.Printf("The sum of %d and %d is %d.\n", num1, num2, num1+num2)
}
