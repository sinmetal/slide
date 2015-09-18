package main

import (
	"fmt"
	"runtime"
)

func helloWorld() {
	fmt.Println("hello world")
}

func main() {
	go helloWorld()

	runtime.Gosched() // goroutine切り替え
}
