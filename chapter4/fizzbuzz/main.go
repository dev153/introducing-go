package main

import "fmt"

func main() {
	num := 1
	for num <= 100 {
		fmt.Print(num)
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if num%5 == 0 {
			fmt.Println("Buzz")
		} else if num%3 == 0 {
			fmt.Println("Fizz")
		}
		num++
	}
}
