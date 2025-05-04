package main

import (
	"fmt"
	"sync"
	"time"
)

func someFunc(wg *sync.WaitGroup, num string,
) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println(num)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go someFunc(&wg, "1")
	go someFunc(&wg, "1")
	go someFunc(&wg, "1")
	fmt.Println("Heloo world")
	wg.Wait()
}
