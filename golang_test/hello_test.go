package main

import "testing"

func TestSayHello(t *testing.T) {
	sentence := sayHello("Adam")

	if sentence != "Hello there Adam\n" {
		t.Error("Fail test:", sentence)
	}
}
