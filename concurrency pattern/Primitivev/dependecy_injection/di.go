package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello,%v", name)
}

func main() {

	//my own prinf func
	Greet(os.Stdout, "harsh")
}
