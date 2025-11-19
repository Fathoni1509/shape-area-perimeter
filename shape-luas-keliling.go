package main

import (
	"fmt"
	"math"
)

// interface shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// untuk bangun datar persegi
type Square struct {
	Side float64
}

// luas
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// keliling
func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// untuk bangun datar persegi panjang
type Rectangle struct {
	Length, Width float64
}

// luas
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// keliling
func (r Rectangle) Perimeter() float64 {
	return 2*r.Length + 2*r.Width
}

// untuk bangun datar segitiga (segitiga sama sisi)
type Triangle struct {
	Base, Height, Side float64
}

// luas
func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

// keliling
func (t Triangle) Perimeter() float64 {
	return 3 * t.Side
}

// untuk bangun datar lingkaran
type Circle struct {
	Radius float64
}

// luas
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// keliling
func (c Circle) Perimeter() float64 {
	return math.Pi * c.Radius * 2
}

// untuk bangun datar trapesium (trapesium siku siku)
type Trapezoid struct {
	SideA, SideB, Height, SideSkew float64
}

// luas
func (t Trapezoid) Area() float64 {
	return (t.SideA + t.SideB) * t.Height / 2
}

// keliling
func (t Trapezoid) Perimeter() float64 {
	return t.SideA + t.SideB + t.Height + t.SideSkew
}

// function process untuk cetak nama, luas, dan keliling
func ProcessShape(sh Shape) {
	var name string // simpan nama shape

	// cocokkan tipe shape yang sesuai
	switch sh.(type) {
	case Square:
		name = "Square"
	case Rectangle:
		name = "Rectangle"
	case Triangle:
		name = "Triangle"
	case Circle:
		name = "Circle"
	case Trapezoid:
		name = "Trapezoid"
	}

	// cetak shape, luas, dan keliling
	fmt.Printf("\nShape: %s | Area: %.2f | Perimeter: %.2f\n\n", name, sh.Area(), sh.Perimeter())

}

// fungsi utama
func main() {
	// persegi
	sq := Square{Side: 4}
	ProcessShape(sq)
	fmt.Println("=========================")

	// persegi panjang
	rec := Rectangle{Width: 4, Length: 3}
	ProcessShape(rec)
	fmt.Println("=========================")

	// segitiga sama sisi
	tri := Triangle{Base: 4, Height: 3, Side: 4}
	ProcessShape(tri)
	fmt.Println("=========================")

	// lingkaran
	cir := Circle{Radius: 10}
	ProcessShape(cir)
	fmt.Println("=========================")

	// trapesium siku siku
	trap := Trapezoid{SideA: 4, SideB: 5, SideSkew: 7, Height: 5}
	ProcessShape(trap)
	fmt.Println("=========================")

}
