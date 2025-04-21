package main

import (
	"fmt"
)

func SumArray(nums []int) (sum int) {

	for _, value := range nums {
		sum += value
	}
	return sum
}

func SumAll(data [][]int) []int {
	lenthofarray := len(data)
	result := make([]int, lenthofarray)
	for i, value := range data {
		result[i] = SumArray(value)
	}

	return result
}

func SumofSlice(data ...[]int) []int {
	var result []int

	for _, value := range data {
		if len(value) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, SumArray(value[1:]))
		}

	}
	return result
}

func main() {
	data := []int{1, 2, 3}
	fmt.Printf(" the slice %v tail  is %v ", data, data[1:])
}
