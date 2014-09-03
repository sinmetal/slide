package main

import "fmt"

func sum(a, b int, c chan<- int) {
	fmt.Println("sum start", a, b)
	c <- a + b
	fmt.Println("sum end", a, b)
}

func main() {
	c := make(chan int) // capacity 0
	go sum(1, 1, c)
	go sum(2, 2, c)
	go sum(3, 3, c) // capacity 0なので、chには値を１つずつしか入れれない

	fmt.Println("wait ch!")
	x := <-c
	y := <-c

	fmt.Println("x y =", x, y)
}
