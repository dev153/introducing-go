package main

import "math"
import "fmt"

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	x, y, width, height float64
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func (rect Rectangle) perimeter() float64 {
	return 2*rect.width + 2*rect.height
}

type Circle struct {
	x, y, r float64
}

func (circ Circle) area() float64 {
	return math.Pi * circ.r * circ.r
}

func (circ Circle) perimeter() float64 {
	return 2 * math.Pi * circ.r
}

type MultiShape struct {
	multiShapes []Shape
}

func (ms *MultiShape) area() float64 {
	result := float64(0)
	for _, aShape := range ms.multiShapes {
		result += aShape.area()
	}
	return result
}

func (ms *MultiShape) perimeter() float64 {
	result := float64(0)
	for _, aShape := range ms.multiShapes {
		result += aShape.perimeter()
	}
	return result
}

func sumAreas(args ...Shape) float64 {
	result := float64(0)
	for _, shape := range args {
		result += shape.area()
	}
	return result
}

func sumPerimeters(args ...Shape) float64 {
	result := float64(0)
	for _, shape := range args {
		result += shape.perimeter()
	}
	return result
}

func main() {
	c1 := Circle{r: 1}
	r1 := Rectangle{width: 1, height: 1}
	fmt.Println(sumAreas(c1, r1))
	fmt.Println(sumPerimeters(c1, r1))
	ms := MultiShape{
		multiShapes: []Shape{
			// Circle{1, 1, 1},
			c1,
			r1,
		},
	}
	fmt.Println(sumAreas(&ms))
	fmt.Println(sumPerimeters(&ms))
}
