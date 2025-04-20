package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("sid")
	want := "Hello,sid"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
