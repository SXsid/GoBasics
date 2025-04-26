package arraylist

type arraylist struct {
	capacity int
	data     []int
}

func NewArraylist(initCapactiy int) *arraylist {
	return &arraylist{capacity: initCapactiy, data: make([]int, 0, initCapactiy)}
}

func (arr *arraylist) persistTheSize() {
	if arr.capacity == len(arr.data) {

		arr.capacity *= 2
		newArray := make([]int, len(arr.data), arr.capacity*2)
		copy(newArray, arr.data)
		arr.data = newArray

	}
}

func (arr *arraylist) add(value int) {
	arr.persistTheSize()
	arr.data = append(arr.data, value)
}

func (arr *arraylist) remove() {
	arr.data[len(arr.data)-1]

}
