package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := int64(1); i < 90; i++ {
		fmt.Println(fib(i))
	}

	end := time.Now()
	fmt.Println(end.UnixNano() - start.UnixNano())
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
