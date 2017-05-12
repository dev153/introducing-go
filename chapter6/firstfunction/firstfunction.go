package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func sum(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total
}

func statistics(xs []float64) (theSum float64, theAverage float64) {
	theSum = sum(xs)
	theAverage = average(xs)
	return
}

func printInts(theNumbers ...int) {
	for _, aNumber := range theNumbers {
		fmt.Println(aNumber)
	}
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(xs)
	fmt.Println(sum(xs))
	fmt.Println(average(xs))
	x, y := statistics(xs)
	fmt.Println(x, y)
	xs = append(xs, 87)
	x, y = statistics(xs)
	fmt.Println(x, y)
	printInts(1, 2, 3)
	printInts(1, 2, 3, 4)
	xi := []int{2, 3, 5, 7, 11, 13}
	printInts(xi...)
}
