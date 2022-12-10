package main

import (
	"fmt"
)

func main() {
	a := 10             //1010
	b := 3              //0011
	fmt.Println(a & b)  //0010 AND
	fmt.Println(a | b)  //1011 OR
	fmt.Println(a ^ b)  //1001 XOR
	fmt.Println(a &^ b) //0100 NOR
	a = 8               //2^3 1000
	fmt.Println(a >> 3) //(2^3)/2^3 = 2^0	0001
	fmt.Println(a << 4) //(2^3)*2^4 = 2^7	10000000
}
