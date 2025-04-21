package main

func SumArray(nums []int) (sum int) {

	for _, value := range nums {
		sum += value
	}
	return sum
}
