package main

import (
	"fmt"
)

func helloWorld() {
	fmt.Println("hello world")
}

func main() {
	go helloWorld() // HL
}
