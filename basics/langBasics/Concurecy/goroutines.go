package main

import (
	"fmt"
	"sync"
	"time"
)

type data struct {
	email        string
	downloadLink string
}

func downloadAndConfrimByEmail(dataSlice []data) {
	var wg sync.WaitGroup
	for i, d := range dataSlice {
		wg.Add(1)
		//concurent / differ it
		go func(d data, index int) {
			defer wg.Done()
			fmt.Printf("donwlaoding the data form %v \n", d.downloadLink)
			time.Sleep(time.Millisecond * 10000)
			fmt.Printf("finishes the %d downlaod", index+1)
		}(d, i)
		fmt.Printf("email has been sent too %v\n", d.email)
	}
	wg.Wait()
}

func main() {
	// downloadAndConfrimByEmail([]data{{email: "dafdsj", downloadLink: "sdfkladslsd"}, {email: "sid@gmail.com", downloadLink: "some random link"}})

	Channel([]int{20, 17, 16})
}
