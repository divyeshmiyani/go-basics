package main

import (
	"fmt"
)

func main() {
	//Concurrency in Go
	job := make(chan int, 100)
	res := make(chan int, 100)
	in := make(chan int, 100)

	go worker(job, res, in)
	//go worker(job, res, in)
	//go worker(job, res, in)
	//go worker(job, res, in)
	//go worker(job, res, in)
	//go worker(job, res, in)
	//go worker(job, res, in)

	for i := 1; i <= 50; i++ {
		job <- i //sends 1 to 50 in job channel
	}
	close(job)
	for m := range res {
		fmt.Printf("%d = %d\n", <-in, m)
	}
}
func worker(job <-chan int, res chan<- int, in chan<- int) {
	for n := range job {
		res <- fibo(n) //calculate nth fibonacci number as send in res channel
		in <- n        //sends n in channel in
	}
	close(res)
	close(in)
}
func fibo(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}
