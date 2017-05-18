package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}() // we need a function call () operator when the function is anonymous

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}() // we need a function call () operator when the function is anonymous

	// Let's see how we can select the channel to read from with the "select" statement...

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
				default:
				fmt.Println("no channels are ready")
			}
		}

	}() // we need a function call () operator when the function is anonymous

	var input string
	fmt.Scanln(&input)
}
