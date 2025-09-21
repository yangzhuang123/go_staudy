package main

import "fmt"

type Shape interface {
	Area()

	Perimeter()
}

type Rectangle struct {
}

func (this *Rectangle) Area() {
	fmt.Println("Rectangle.Area....")
}

func (this *Rectangle) Perimeter() {
	fmt.Println("Rectangle.Perimeter....")
}

type Circle struct {
}

func (this *Circle) Area() {
	fmt.Println("Circle.Area....")
}

func (this *Circle) Perimeter() {
	fmt.Println("Circle.Perimeter....")
}

func main() {
	var shape Shape
	shape = &Rectangle{}

	shape.Area()
	shape.Perimeter()

	shape = &Circle{}
	shape.Area()
	shape.Perimeter()
}
