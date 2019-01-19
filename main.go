package main

import (
	"fmt"
	"math"
	"strconv"
)

//Shape shape
type Shape interface {
	Square() float64
	Perim() float64
	Say()
}

// ShapeList shape list
type ShapeList struct {
	List map[string]Shape
}

// Rectangle struct with rectangle data
type Rectangle struct {
	Width  float64
	Height float64
	Type   string
}

// Triangle struct with triangle data
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
	Type  string
}

// Circle struct with circle data
type Circle struct {
	Radius float64
	Type   string
}

// Square method calculate square of the rectangle
func (r Rectangle) Square() float64 {
	return float64(r.Height * r.Width)
}

// Perim method calculate perimiter of the rectangle
func (r Rectangle) Perim() float64 {
	return float64(r.Width*2 + r.Height*2)
}

// Say method print all data of the rectangle
func (r Rectangle) Say() {
	fmt.Printf("%s Width: %v Height: %v Square: %v Perim: %v \n",
		r.Type, r.Width, r.Height, r.Square(), r.Perim())
}

// Square method calculate square of the Triangle
func (r Triangle) Square() float64 {
	if (r.SideA > r.SideB) && (r.SideA > r.SideC) {
		return float64(r.SideB * r.SideC / 2)
	}

	if (r.SideB > r.SideA) && (r.SideB > r.SideC) {
		return float64(r.SideA * r.SideC / 2)
	}

	return float64(r.SideA * r.SideB / 2)
}

// Perim method calculate perimiter of the Triangle
func (r Triangle) Perim() float64 {
	return float64(r.SideA + r.SideB + r.SideC)
}

// Say method print all data of the Triangle
func (r Triangle) Say() {
	fmt.Printf("%s SideA: %v SideB: %v SideC: %v Square: %v Perim: %v \n",
		r.Type, r.SideA, r.SideB, r.SideC, r.Square(), r.Perim())
}

// Square method calculate square of the Circle
func (r Circle) Square() float64 {
	return float64(r.Radius * r.Radius * math.Pi)
}

// Perim method calculate perimiter of the Circle
func (r Circle) Perim() float64 {
	return float64(r.Radius * 2 * math.Pi)
}

// Say method print all data of the Circle
func (r Circle) Say() {
	fmt.Printf("%s Radius:%v Square: %v Perim: %v \n",
		r.Type, r.Radius, r.Square(), r.Perim())
}

// Say method for shape list
func (l ShapeList) Say() {
	for _, shape := range l.List {
		shape.Say()
	}
}

// Add shape to ShapeList
func (l *ShapeList) Add(s Shape) {
	index := len(l.List)
	l.List[strconv.Itoa(index)] = s
}

func newRectangle(Width, Height float64) Rectangle {
	return Rectangle{Width, Height, "Rectangle"}
}
func newTriangle(SideA, SideB, SideC float64) Triangle {
	return Triangle{SideA, SideB, SideC, "Triangle"}
}
func newCircle(Radius float64) Circle {
	return Circle{Radius, "Circle"}
}
func main() {
	list := ShapeList{List: make(map[string]Shape)}

	list.Add(newCircle(5))
	list.Add(newRectangle(5, 6))
	list.Add(newTriangle(3, 4, 5))

	list.Say()
}
