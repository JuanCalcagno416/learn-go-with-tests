package main

import "math"

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Perimeter will return the perimeter of a Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Perimeter will return the perimeter of a Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// Perimeter will return the perimeter of a Triangle
func (t Triangle) Perimeter() float64 {
	h := math.Sqrt(t.Base*t.Base + t.Height*t.Height)
	return t.Height + t.Base + h
}

// Area will return the Area a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area will return the Area a Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Area will return the Area a Triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
