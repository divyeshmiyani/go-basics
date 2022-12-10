package main

import (
	"fmt"
)

func main() {
	// 	grades := [...]int{66, 40, 54} // `...` takes size of array automatically here 3
	// 	fmt.Printf("Grades = %v", grades)

	a := [...]int{1, 2, 3} //array

	// a copied to b
	b := a // all values copied
	// b := &a // address of values copied

	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("Length of a = ", len(a))
	fmt.Println("Length of b = ", len(b))
}
