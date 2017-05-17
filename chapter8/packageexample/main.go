package main

import "fmt"
import mymath "github.com/dev153/introducing-go/chapter8/packageexample/math"
import "math"

func main() {
	fmt.Println("We will test the package mechanism of Go.")
	xs := []float64{1, 2, 3, 4}
	avg := mymath.Average(xs)
	fmt.Println(avg)
	fmt.Println(math.Phi)
}
