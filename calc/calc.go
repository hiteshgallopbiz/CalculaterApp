package main

import (
	"fmt"
)

func main() {

	var num1, num2 int
	var choice string

	fmt.Println("Enter first num")
	fmt.Scan(&num1)

	fmt.Println("Enter your operater")
	fmt.Scan(&choice)

	fmt.Println("Enter second num")
	fmt.Scan(&num2)

	switch choice {
	case "+":
		fmt.Printf("Addition: %d \n", Addition(num1, num2))
	case "-":
		fmt.Printf("Difference: %d \n", Substraction(num1, num2))
	case "*":
		fmt.Printf("Product: %d \n", Multiplication(num1, num2))
	case "/":
		fmt.Printf("Division: %d \n", Division(num1, num2))
	default:
		fmt.Println("Wrong option try with +, -, / and *")
	}

}

func Addition(num1, num2 int) int {
	result := num1 + num2
	return result

}

func Substraction(num1, num2 int) int {
	return num1 - num2

}

func Multiplication(num1, num2 int) int {
	return num1 * num2

}

func Division(num1, num2 int) int {
	return num1 / num2

}
