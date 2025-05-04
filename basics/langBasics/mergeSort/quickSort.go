package mergesort

//the partion approch

func partion(arr []int, low, high int) int {
	pivot := arr[high]

	startPoint := low
	//all samll elemt in the starting of the sub array
	for i := low; i < high; i++ {
		if arr[i] < pivot {
			arr[i], arr[startPoint] = arr[startPoint], arr[i]
			startPoint++
		}

	}
	//now just after all small palce the pivot
	arr[high], arr[startPoint] = arr[startPoint], arr[high]
	return startPoint
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partion(arr, low, high)

		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}

}
