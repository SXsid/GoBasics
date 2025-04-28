package main

import (
	"errors"
	"fmt"
)

type user struct {
	name        string
	phoneNumber int
}

func convertSliceToMap(nameSlice []string, phoneNumberSlice []int) (map[string]user, error) {
	if len(nameSlice) != len(phoneNumberSlice) {
		return nil, errors.New("invalid sizes")
	}

	requiredMap := make(map[string]user)

	for i := 0; i < len(nameSlice); i++ {
		requiredMap[nameSlice[i]] = user{
			name:        nameSlice[i],
			phoneNumber: phoneNumberSlice[i],
		}
	}

	return requiredMap, nil

}

func maMap() {

	firstMap, err := convertSliceToMap([]string{"aman", "harsh"}, []int{1234, 7890839})

	if err != nil {
		fmt.Errorf(err.Error())
	}
	fmt.Println(firstMap)
	if value1, ok := firstMap["aman"]; !ok {
		fmt.Printf("no value exist")
	} else {
		fmt.Println(value1)
	}
	firstMap["sas"] = user{
		name:        "sas",
		phoneNumber: 9352452,
	}
	if _, ok := firstMap["sas"]; !ok {
		fmt.Printf("no value exist")
	} else {
		delete(firstMap, "sas")
		fmt.Println("done")
	}

}
