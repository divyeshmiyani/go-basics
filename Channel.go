package main

import "fmt"

func main() {
	c := make(chan string, 2) //second argument is capacity of channel
	c <- "hello"
	c <- "world"
	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)

}
