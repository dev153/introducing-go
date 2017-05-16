package main

import "fmt"

func sum(sl []int) int {
	result := 0
	for _, value := range sl {
		result += value
	}
	return result
}

func main() {
	var sl []int
	sl = append(sl, 1, 2, 3, 4)
	result := sum(sl)
	fmt.Println(result)
}
