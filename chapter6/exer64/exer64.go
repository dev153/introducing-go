package main

import "fmt"

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (oddNumber uint) {
		oddNumber = i
		i += 2
		return
	}
}

func main() {
	oddGenerator := makeOddGenerator()
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())
	fmt.Println(oddGenerator())
}
