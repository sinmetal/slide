package main

import "fmt"

func main() {
	i := func(x int, y int) int {
		return x + y
	}(2, 3)

	fmt.Println(i, "")
}
