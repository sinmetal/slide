package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	count := int64(50)
	c := make(chan int64)

	for i := int64(1); i < count; i++ {
		go fibonacci(i, c)
	}

	for i := int64(1); i < count; i++ {
		ret := <-c
		fmt.Println(ret)
	}

	end := time.Now()
	fmt.Println(end.UnixNano() - start.UnixNano())
}

func fibonacci(n int64, c chan int64) {
	c <- fib(n)
}

func fib(n int64) int64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}

	return (fib(n-2) + fib(n-1))
}
