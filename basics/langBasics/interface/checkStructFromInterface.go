package main

import (
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

type shape interface {
	getArea() float64
}

func (c circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (rect rectangle) getArea() float64 {
	return rect.height * rect.width
}

func getShapeInfo(s shape) (string, float64) {
	switch v := s.(type) {
	case circle:
		return "circle", v.getArea()
	case rectangle:
		return "rectangle", v.getArea()
	default:
		return "unknown", 0
	}
}

func main() {

}
