package main

import "fmt"

func main() {
	fmt.Println("Inside main function")
	func() {
		fmt.Println("Inside inline function")
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		fmt.Println("inline function ends")
	}()
	fmt.Println("main function ends")
}
