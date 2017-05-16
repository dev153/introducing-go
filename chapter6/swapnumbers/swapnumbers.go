package main

import (
	"fmt"
)

func swap(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func main() {
	x := 2
	y := 1
	fmt.Printf("Numbers before swap x=%d y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("Numbers after swap x=%d y=%d\n", x, y)
}
