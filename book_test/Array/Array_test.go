package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumArray(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSummAll(t *testing.T) {

	sliceOfSlice := [][]int{{1, 2}, {3, 4}}

	got := SumAll(sliceOfSlice)
	want := []int{3, 7}

	//it's not type safe
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	//anonymus fuction
	commonErrorMsg := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumofSlice([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		commonErrorMsg(t, got, want)

	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumofSlice([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		commonErrorMsg(t, got, want)
	})

}
