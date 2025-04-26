package maps

import "testing"

func TestDictonry(t *testing.T) {
	dict := Dict{"key1": "value1"}

	got := dict.Search("key1")
	want := "value1"
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
