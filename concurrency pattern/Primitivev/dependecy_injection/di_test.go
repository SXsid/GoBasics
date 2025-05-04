package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	//my own std out
	write := bytes.Buffer{}
	//print f make the stirnginto buffer and tell to print on the std out we are telling updated our buffer instead
	Greet(&write, "sid")
	got := write.String()
	want := "Hello,sid"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
