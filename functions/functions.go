package functions

import "fmt"

func Addition(num1, num2 int) int {
	return num1 + num2
}

func Subtraction(num1, num2 int) int {
	return num1 - num2
}

func Division(num1, num2 int) {
	var result int
	if num2 == 0 {
		fmt.Println("cannot divide by zero")
		return
	}
	result = num1 / num2
	fmt.Println(result)
}

func Multiplication(num1, num2 int) int {
	return num1 * num2
}

func Modular(num1, num2 int) int {
	return num1 % num2
}
