package httpsample

import (
	"io"
	"log"
	"net/http"
)

// hello golang, the web server
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, golang!\n")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
