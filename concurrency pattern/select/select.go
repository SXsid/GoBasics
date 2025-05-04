package main

import (
	"fmt"
)

func main() {

	firstChaneel := make(chan string)
	secondChanel := make(chan string)

	go func() {

		firstChaneel <- "channel 1"
	}()

	go func() {

		secondChanel <- "channel 2"
	}()

	select {
	case first := <-firstChaneel:
		fmt.Println("first completed first", first)
	case second := <-secondChanel:
		fmt.Println("done with second", second)
	}
}
