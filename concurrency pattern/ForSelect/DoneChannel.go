package main

import "fmt"

//for is to work infiltey  to  insert or read data
//select allow the parent to control the flow of execution
//hence we always warap for with select to control the channel execution
// func main() {
// 	channel := make(chan string)
// 	go doSomething(channel)
// 	for {
// 		data := <-channel
// 		fmt.Println(data)
// 	}
// }

// func doSomething(channel chan string) {
// 	for {
// 		channel <- "we can do it "
// 	}
// }

//to control this we need done chanel as if some task which are done parent will stop the sender routing
func main() {
	channel := make(chan string, 3)
	done := make(chan bool)
	count := 0
	go doSomething(channel, done)
	//we will only read first 3 vaue and close the cahnel
	for count < 3 {
		data := <-channel
		fmt.Println(data)
		count++
	}
	//control the rouinge form parent
	done <- true

	for count < 5 {
		fmt.Println("main fction is still goin")
		count++
	}

}

func doSomething(channel chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			close(channel)
			return
		case channel <- "do something":

		}

	}

}
