package main

import "fmt"

func printint(xptr *int) {
	fmt.Println(*xptr) // 2, 3
	*xptr++
}

func main() {
	xptr := new(int)
	x := 2
	xptr = &x
	fmt.Println(*xptr) // 2
	fmt.Println(x)     // 2
	printint(xptr)
	printint(&x)
	fmt.Println(x) // 4
}
