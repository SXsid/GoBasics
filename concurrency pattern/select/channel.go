package main

import (
	"fmt"
	"time"
)

// (excution time:3sec)
// treting the chanel basd as normal call / making is sync not parallel
func Channel() {
	firstChaneel := make(chan string)
	secondChanel := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		firstChaneel <- "first channel data"
	}()
	data := <-firstChaneel
	fmt.Println(data)
	go func() {
		time.Sleep(1 * time.Second)
		secondChanel <- 2
	}()

	secodnData := <-secondChanel
	fmt.Println(data, secodnData)
}

// now it parralel but result are still sequental (excution time:2sec)
// even tought the seond get it's job but until we complet the first the result will be hunged in buffer zone
func ChannelSync() {
	firstChaneel := make(chan string)
	secondChanel := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		firstChaneel <- "channel data"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		secondChanel <- 2
	}()
	data := <-firstChaneel
	secodnData := <-secondChanel
	fmt.Println(secodnData, data)
}
