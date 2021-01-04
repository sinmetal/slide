package main

import "fmt"

func sum(a, b int, c chan<- int) {
	fmt.Println(a, b)
	c <- a + b
}

func main() {
	c := make(chan int)
	go sum(3, 7, c) // 1つしかchに書き込みしない

	fmt.Println("wait ch!")
	x, y := <-c, <-c // deadlock!!

	fmt.Println(x, y)
}
