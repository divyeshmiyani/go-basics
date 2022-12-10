package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3} //slice
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}
