package mergesort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestMergeSort(t *testing.T) {
	t.Run("SmallCorrectnessCase", func(t *testing.T) {
		nums := []int{5, 4, 31, 3, 2, 1, -1}
		mergesort(nums, 0, len(nums)-1)
		want := []int{-1, 1, 2, 3, 4, 5, 31}
		if !reflect.DeepEqual(nums, want) {
			t.Errorf("wanted %v got %v", want, nums)
		}
	})

	t.Run("Benchmark1000", func(t *testing.T) {
		size := 1000_000_000
		randNums := make([]int, size)
		for i := range randNums {
			randNums[i] = rand.Intn(10000)
		}

		start := time.Now()
		mergesort(randNums, 0, len(randNums)-1)
		duration := time.Since(start)
		fmt.Printf("Time taken to sort %d elements: %v", size, duration)

		for i := 1; i < len(randNums); i++ {
			if randNums[i] < randNums[i-1] {
				t.Errorf("array is not sorted on %d index,", i-1)
			}
		}
	})
}

func TestQuickSort(t *testing.T) {
	nums := []int{5, 4, 31, 3, 2, 1, -1}
	quickSort(nums, 0, len(nums)-1)
	want := []int{-1, 1, 2, 3, 4, 5, 31}
	if !reflect.DeepEqual(nums, want) {
		t.Errorf("wanted %v got %v", want, nums)
	}
}
