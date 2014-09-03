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

func handler(w http.ResponseWriter, r *http.Request) {
	tasks := [...]*task{
		&task{"https://dl.dropboxusercontent.com/u/15506977/qiitaitem.json", make(chan string)},
		&task{"https://dl.dropboxusercontent.com/u/15506977/qiitaitem-java.json", make(chan string)},
	}

	for _, t := range tasks {
		go func(t *task) {
			qiita(t)
		}(t)
	}

	w.Header().Set("Content-Type", "text")
	w.Header().Set("charset", "utf-8")
	for _, t := range tasks {
		fmt.Fprintln(w, <-t.result)
	}
}

func qiita(t *task) {
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
	http.HandleFunc("/", handler)
	http.ListenAndServe(":12345", nil)
}
