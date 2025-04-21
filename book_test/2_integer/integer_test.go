package integer

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	returnErrormsg(t, expected, sum)
}

func returnErrormsg(t testing.TB, sum, expected int) {
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
