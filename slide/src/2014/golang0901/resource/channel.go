package main

import "fmt"

func sum(a, b int, c chan<- int) {
	fmt.Println(a, b)
	c <- a + b
}

func main() {
	c := make(chan int)
	go sum(3, 7, c)
	go sum(4, 8, c)

	fmt.Println("wait ch!")
	x, y := <-c, <-c // chから2つの値を受け取るまでblock

	fmt.Println(x, y)
}
