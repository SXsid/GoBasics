package main

import (
	"bytes"
	"testing"
)

func TestCoundown(t *testing.T) {
	buffer := &bytes.Buffer{}
	countDown(buffer)
	got := buffer.String()
	want := `3
2
1
Go!`
	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

}
