package stack

import (
	"testing"
)

func TestStack(t *testing.T) {

	Errormsg := func(t testing.TB, s *Stack, want int) {
		t.Helper()
		get := s.Peek()
		if get != want {
			t.Errorf("got %d wanted %d", get, want)
		}

	}
	t.Run("stack push", func(t *testing.T) {
		stack := &Stack{}

		stack.push(10)

		want := 10
		Errormsg(t, stack, want)

	})

	t.Run("stack pop", func(t *testing.T) {
		stack := &Stack{}
		stack.push(30)
		stack.push(20)
		stack.pop()
		Errormsg(t, stack, 30)

	})
}
