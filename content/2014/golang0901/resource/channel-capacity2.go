package main

import "fmt"

func sum(a, b int, c chan<- int) {
	fmt.Println("sum start", a, b)
	c <- a + b
	fmt.Println("sum end", a, b)
}

func main() {
	c := make(chan int, 5) // capacity 5
	go sum(1, 1, c)
	go sum(2, 2, c)
	go sum(3, 3, c) // capacity 5なので、3つのgoroutineが値を連続で入れることができる

	fmt.Println("wait ch!")
	x := <-c
	y := <-c

	fmt.Println("x y =", x, y)
}
