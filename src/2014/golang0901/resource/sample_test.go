package main

import "testing"

func TestSample(t *testing.T) {
	if "hello" != "world" {
		t.Fatalf("Non-expected text : %v", "hello world")
	}
}
