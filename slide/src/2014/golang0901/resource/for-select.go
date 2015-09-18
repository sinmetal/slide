package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			c <- "hello!"
		}
	}()

	for {
		select {
		case msg := <-c:
			fmt.Println("received", msg)
		}
	}
}
