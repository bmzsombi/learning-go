package structsinterfaces

import "math"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func NewCircle(Radius float64) Circle {
	return Circle{Radius: Radius}
}

func NewRectangle(Width, Height float64) Rectangle {
	return Rectangle{Width: Width, Height: Height}
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}
