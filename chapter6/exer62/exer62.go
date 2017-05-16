package main

import "fmt"

func half(n int) (int, bool) {
	res1 := 0
	res2 := false
	if n%2 == 0 {
		res1 = 1
		res2 = true
	}
	return res1, res2
}

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
}
