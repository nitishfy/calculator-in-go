package main

import (
	"calculator-in-go/functions"
	"fmt"
)

func main() {
	var num1 int

	var num2 int

	var operator rune

	for { // infinite loop

		fmt.Print("Enter the first number: ")

		fmt.Scanln(&num1)

		fmt.Print("Enter the second number: ")

		fmt.Scanln(&num2)

		fmt.Print("Enter the operator(Press x to quit): ")

		fmt.Scanf("%c\n", &operator) // taking user input

		if operator == '+' || operator == '-' || operator == '/' || operator == '*' || operator == '%' {

			if operator == '+' {
				fmt.Print("The result is: ")
				fmt.Println(functions.Addition(num1, num2))
			}

			if operator == '-' {
				fmt.Print("The result is:")
				fmt.Println(functions.Subtraction(num1, num2))
			}

			if operator == '/' {
				fmt.Print("The result is: ")
				functions.Division(num1, num2)
			}

			if operator == '*' {
				fmt.Print("The result is:")
				functions.Multiplication(num1, num2)
			}

			if operator == '%' {
				fmt.Print("The result is:")
				functions.Modular(num1, num2)
			}

		} else if operator == 'x' || operator == 'X' {
			return

		} else {
			fmt.Println("Ivalid operator")
		}
	}

}
