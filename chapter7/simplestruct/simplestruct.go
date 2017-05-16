package main

import "fmt"
import "math"

type Circle struct {
	x, y, r float64
}

func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2.0 * math.Pi * c.r
}

func (c *Circle) multiplyPerimeterByFactor(f float64) float64 {
	return c.perimeter() * f
}

func main() {
	var c1 Circle           // initialization of all fields with zero values.
	c2 := Circle{0, 0, 4}   // initialization of fields with given values.
	c3 := &Circle{0, 0, 10} // initialization of a pointer to a structure with given values
	c4 := Circle{}
	c5 := new(Circle)
	c6 := Circle{y: 2, x: 1, r: 3}
	c7 := Circle{x: 4, y: 4}

	fmt.Println(c1, c2, *c3, c4, c5, c6, c7)
	fmt.Println(circleArea(c6))
	fmt.Println(c6.perimeter())
	fmt.Println(c6.multiplyPerimeterByFactor(2))
}
