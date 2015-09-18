package main

import (
	"fmt"
	"time"
	"strconv"
)

func makeBook(c chan <- string) {
	var i int = 1
	for {
		time.Sleep(2 * time.Second)
		b := "book" + strconv.Itoa(i)
		fmt.Println("make book! = ", b)
		c <- b
		i++
	}
}

func moveBook(book string, c2 chan <- string) {
	status := [4]string{"", "", "", ""}
	for i := 0; i < 4; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("move book", status, "->", book)
		status[i] = "."
	}
	c2 <- book
}

func fireBook(c <- chan string) {
	for {
		select {
		case book := <-c:
			fmt.Println("fire !! -> ", book)
		}
	}
}

func main() {
	c1 := make(chan string, 5)
	c2 := make(chan string, 5)

	go makeBook(c1)
	go fireBook(c2)

	for {
		select {
		case book := <-c1:
			go moveBook(book, c2)
		}
		
		
	}
}