package main

import "fmt"

func fibo(n uint) uint {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-2) + fibo(n-1)
	}
}

func main() {
	for i := uint(0); i < 10; i++ {
		fmt.Println("i:", i, "fib(i):", fibo(i))
	}
}
