package main

import "fmt"

var i string = "String initialized outside main() function. Is it visible inside main()?"

func main() {
	var x string = "Hello, "
	fmt.Println(x)
	var y string
	y = "World"
	fmt.Println(y)
	x = x + y
	fmt.Println(x)
	x += " again..."
	fmt.Println(x)
	z := "This is a string...its type is inferred by Go..."
	fmt.Println(z)
	fmt.Println(i)

	const j string = "A constant string..."
	fmt.Println(j)
	// j = "This one is not feasible...it won't compile..."
}
