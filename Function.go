package main

import (
	"fmt"
)

func Addition(num1, num2 int) int {
	out := num1 + num2
	return out
}

func Subtraction(num1, num2 int) (out int) { //out is decleared here
	out = num1 - num2
	return // returns out
}

func Multiplication(num1, num2 int) (out int) {
	out = num1 * num2
	return out
}

func calc(num1, num2 int) (add, sub, mul int) {
	add = num1 + num2
	sub = num1 - num2
	mul = num1 * num2
	return //returns add, sub, mul
}

func main() {
	var a, b = 10, 1
	fmt.Println(Addition(a, b))
	fmt.Println(Subtraction(a, b))
	fmt.Println(Multiplication(a, b))
	fmt.Println()
	fmt.Println(calc(a, b))
}
