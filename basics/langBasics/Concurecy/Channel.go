package main

import "fmt"

func Channel(ageSlice []int) {
	isAdultChan := make(chan bool)
	go func() {
		for _, age := range ageSlice {
			if age > 18 {
				isAdultChan <- true
				continue
			}
			isAdultChan <- false
		}
		defer close(isAdultChan)
	}()

	for isAdult := range isAdultChan {
		fmt.Printf("is the perosn is adlut=> %t", isAdult)
	}
}
