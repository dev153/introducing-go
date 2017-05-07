package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println(total / float64(len(y)))

	total = 0
	// Special for loop for iterating arrays.
	for _, value := range y {
		total += value
	}
	fmt.Println(total / float64(len(y)))

	z := [5]float64{98, 93, 77, 82, 83}

	total = 0
	for _, value := range z {
		total += value
	}

	fmt.Println(total / float64(len(z)))
}
