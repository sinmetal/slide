package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type task struct {
	url    string
	result chan string
}

func fetch(t *task) {
	res, err := http.Get(t.url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	t.result <- string(body)
}

func main() {
	tasks := [...]*task{
		&task{"https://dl.dropboxusercontent.com/u/15506977/helloworld.txt", make(chan string)},
		&task{"https://dl.dropboxusercontent.com/u/15506977/hellogolang.txt", make(chan string)},
	}

	for _, t := range tasks {
		go func(t *task) {
			fetch(t)
		}(t)
	}

	for _, t := range tasks {
		fmt.Println(<-t.result)
	}
}
