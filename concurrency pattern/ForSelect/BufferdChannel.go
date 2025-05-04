package main

import (
	"fmt"
	"time"
)

func BufferdChannel() {
	myChannel := make(chan string, 3)

	go func() {
		chars := []string{"a", "b"}
		for _, value := range chars {
			myChannel <- value
		}

	}()
	time.Sleep(100 * time.Millisecond)
	for {
		select {
		case val := <-myChannel:
			fmt.Println("Received:", val)
		default:
			fmt.Println("No more data in channel right now.")
			//kill the loop if data over
			return
		}
	}

}
