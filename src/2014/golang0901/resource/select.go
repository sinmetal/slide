package main

import (
	"fmt"
	"time"
)

func message(m string, sec float32, c chan string) {
	time.Sleep(time.Duration(sec) * time.Second)
	c <- m
}

// start
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go message("two", 2, c2)
	go message("one", 1, c1)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println("received", msg)
		case msg := <-c2:
			fmt.Println("received", msg)
		}
	}
}

// end
