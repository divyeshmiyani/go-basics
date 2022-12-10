package main

import (
	"fmt"
)

func Addition(num1 int, num2 int) {
	result := num1 + num2
	fmt.Println(result)
}
func Subtraction(num1, num2 int) {
	result := num1 - num2
	fmt.Println(result)
}

func main() {
	var a, b = 10, 1
	Addition(a, b)
	Subtraction(a, b)
}
