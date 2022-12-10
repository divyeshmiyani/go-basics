package main

import (
	"fmt"
)

func main() {
	var a bool = true // default value of bool variable is false
	fmt.Printf("%v %T\n", a, a)
	var b = 5 //default value of int variable is 0
	fmt.Printf("%v %T\n", b, b)
	var comp = 1 + 2i //default is complex128complex(1,2)
	fmt.Printf("%v %T\n", comp, comp)
	fmt.Printf("real part=%v %T\n", real(comp), real(comp))
	fmt.Printf("imaginary part=%v %T\n", imag(comp), imag(comp))
}
