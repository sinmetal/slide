package httpsample

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// start
func TestHelloServer(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello:12345", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	HelloHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Non-expected status code: %v\nbody: %v", w.Code, w.Body)
	}

	if w.Body.String() != "hello, golang!\n" {
		t.Fatalf("Non-expected body : %v", w.Body)
	}
}

// end
