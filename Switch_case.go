package main

import (
	"fmt"
)

func main() {
	num := 5

	switch num {
	case 1:
		fmt.Println("ONE")
	case 2:
		fmt.Println("TWO")
	case 3:
		fmt.Println("THREE")
	default:
		fmt.Println("NONE")
	}
}
