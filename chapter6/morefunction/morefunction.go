package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}
	z := add(1, 2)
	fmt.Println(z)

	x := 0

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
