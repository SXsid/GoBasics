package main

import "fmt"

type customStruct struct {
	name string
}

func (cust customStruct) Error() string {
	return fmt.Sprintf("the error occuref for%v ", cust.name)
}

func main() {
	fmt.Println(customStruct{name: "sud"}.Error())

}
