package main

import "fmt"

func createDict(name []string) map[string]map[string]int {
	dict := make(map[string]map[string]int)
	for _, value := range name {
		firstLetter := value[0:1]
		if dict[firstLetter] == nil {
			dict[firstLetter] = make(map[string]int)
		}
		dict[firstLetter][value]++
	}
	return dict
}

func main() {

	fmt.Println(createDict([]string{"aman", "aman", "amit", "harhs"}))

}
