package test

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	dataToTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name:    "Reactangle",
			shape:   Rectangle{12, 6},
			hasArea: 72.0,
		}, {
			name:    "circle",
			shape:   Circle{10},
			hasArea: 314.1592653589793,
		}, {
			name:    "Triangle",
			shape:   Triangle{12, 6},
			hasArea: 36.0,
		},
	}

	for _, data := range dataToTest {
		got := data.shape.Area()
		if got != data.hasArea {
			t.Errorf("%#v  got %g hasArea %g", data.shape, got, data.hasArea)
		}
	}
}
