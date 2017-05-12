package main

import "fmt"

func first() {
	fmt.Println("Inside the first() function...")
}

func second() {
	fmt.Println("Inside the second() fuction...")
}

func main() {
	defer second()
	first()
	fmt.Println("exiting the main() method...")
}
