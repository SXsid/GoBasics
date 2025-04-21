package test

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	Base   float64
	Height float64
}

// defineing a fuction whihc is bind with this struct like call wtiht the object in js
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (tri Triangle) Area() float64 {
	return (tri.Base * tri.Height) / 2
}
