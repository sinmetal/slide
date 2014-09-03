package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("hello!!")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(i)
}
