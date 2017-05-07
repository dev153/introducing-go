package main

import "fmt"

func main() {
	x := make([]float64, 5, 10)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x)

	y := [5]float64{1, 2, 3, 4, 5}
	fmt.Println(y)
	z := y[0:5]
	fmt.Println(z)
	fmt.Println(z[0:5])
	fmt.Println(z[1:4])
	fmt.Println(z[2:3])
	fmt.Println(z[1:])
	fmt.Println(z[:4])
	fmt.Println(z[:])

	// a1 := make([]float64, 0)
	var a1 []float64 // an empty slice
	fmt.Println(a1)
	a1 = append(a1, 0)
	fmt.Println(a1)
	a1 = append(a1, 1, 2)
	fmt.Println(a1)
	a2 := []float64{} // another representation of an empty slice
	fmt.Println(a2)

	b1 := []int{1, 2, 3, 4, 5}
	// b2 := []int{}
	b2 := make([]int, 2)
	copy(b2, b1)
	fmt.Println(b1, b2)
}
