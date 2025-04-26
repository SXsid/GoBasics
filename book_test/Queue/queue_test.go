package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {

	Errormsg := func(t testing.TB, queue *Queue, want int) {
		t.Helper()
		get := queue.Peek()
		if get != want {
			t.Errorf("got %d wanted %d", get, want)
		}

	}
	t.Run("enqueue", func(t *testing.T) {
		queue := &Queue{}

		queue.Enqueue(10)

		want := 10
		Errormsg(t, queue, want)

	})

	t.Run("dequeu", func(t *testing.T) {
		queue := &Queue{}
		queue.Enqueue(30)
		queue.Enqueue(20)
		queue.Dqueue()
		Errormsg(t, queue, 20)

	})
}
