package main

import "fmt"

//steps

//we ahve 4 emial
//we nned to fetch data  for 4
//then callucte the anlytics for 4 of the

//it's a have fuction
//in js  i have used promisall to fetch adn prmise to calucte for concurent exectui
//in how we can achive paramely wiht go rouinn using pipeline
//how pipeline work in js even tough we fetched the data cocuretn but we can't move forward with the next step if all the compel
//here we will pass data as soon it ened thorugh channel and keep doing till chanel done/ close
//our main and alytic will be connected to a chnel to read data to
func main() {
	emailSlice := []string{"random1@gamil.com", "random1@gamil.com", "random1@gamil.com", "random1@gamil.com"}

	fetchedDataChaneel := getData(emailSlice)
	analyticsChannel := processData(fetchedDataChaneel)
	for _value := range analyticsChannel {

		fmt.Println(_value)
	}

}

func getData(emailSlice []string) <-chan string {
	dataChannel := make(chan string)
	//if we don't want data from child go routine we can just spin it and let it comple the work and close by self
	go func() {
		for _, value := range emailSlice {
			//to db calls
			fmt.Println(value)
			dataChannel <- fmt.Sprintf("this is data for the %v", value)
		}
		close(dataChannel)
	}()
	return dataChannel
}

func processData(dataChannel <-chan string) (result <-chan interface{}) {
	anylicticCheenl := make(chan interface{})
	//if we don't want data from child go routine we can just spin it and let it comple the work and close by self
	go func() {
		for value := range dataChannel {
			//to db calls
			fmt.Println(value)
			anylicticCheenl <- fmt.Sprintf("this is anyltics data for the %v", value)
		}
		close(anylicticCheenl)
	}()
	return anylicticCheenl

}
