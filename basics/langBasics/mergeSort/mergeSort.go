package mergesort

import (
	"sync"
)

const threshold = 100_000

func mergesort(arr []int, low int, high int) {
	if low >= high {
		return
	}

	mid := low + (high-low)/2

	var wg sync.WaitGroup
	if high-low > threshold {
		wg.Add(2)
		go func() {
			defer wg.Done()
			mergesort(arr, low, mid)
		}()
		go func() {
			defer wg.Done()
			mergesort(arr, mid+1, high)
		}()
		wg.Wait()
	} else {
		mergesort(arr, low, mid)
		mergesort(arr, mid+1, high)
	}

	// Merge step
	tmp := make([]int, 0, high-low+1)
	i, j := low, mid+1
	for i <= mid && j <= high {
		if arr[i] <= arr[j] {
			tmp = append(tmp, arr[i])
			i++
		} else {
			tmp = append(tmp, arr[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, arr[i])
		i++
	}
	for j <= high {
		tmp = append(tmp, arr[j])
		j++
	}
	copy(arr[low:], tmp)
}
