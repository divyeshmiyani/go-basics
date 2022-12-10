package main

import (
	"fmt"
)

const (
	isAdmin            = 1 << iota //00000001
	isHeadquarters                 //00000010
	canSeeFinancials               //00000100
	canSeeAfrica                   //00001000
	canSeeAsia                     //00010000
	canSeeEurope                   //00100000
	canSeeNorthAmerica             //01000000
	canSeeSouthAmerica             //10000000
)

func main() {
	var person byte = canSeeAfrica | canSeeAsia | canSeeFinancials | canSeeSouthAmerica //10011101
	fmt.Printf("person %b\n", person)
	if person&isAdmin == isAdmin { //here, person & isAdmin = 00000001
		fmt.Printf("Person is Admin")
	} else {
		fmt.Printf("Person is not Admin")
	}
}
