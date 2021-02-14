package main

import (
	"fmt"
	"testing"
)

func checkPerimeter(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Perimeter()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
func TestPerimeter(t *testing.T) {
	areaCases := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 36.0},
		{shape: Circle{Radius: 10}, want: 62.83185307179586},
		{shape: Triangle{Base: 12, Height: 6}, want: 31.41640786499874},
	}
	for _, c := range areaCases {
		checkPerimeter(t, c.shape, c.want)
	}
}

func checkArea(t testing.TB, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
func TestArea(t *testing.T) {

	areaCases := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, c := range areaCases {
		checkArea(t, c.shape, c.want)
	}

}
func ExampleArea() {
	r := Rectangle{Width: 12, Height: 6}
	a := r.Area()
	fmt.Println(a)
	// Output: 72
}

func ExamplePerimeter() {
	r := Rectangle{Width: 12, Height: 6}
	a := r.Perimeter()
	fmt.Println(a)
	// Output: 36
}

func BenchmarkArea(b *testing.B) {
	r := Rectangle{Width: 12, Height: 6}
	for i := 0; i < b.N; i++ {
		r.Area()
	}
}

func BenchmarkPerimeter(b *testing.B) {
	r := Rectangle{Width: 12, Height: 6}
	for i := 0; i < b.N; i++ {
		r.Perimeter()
	}
}
