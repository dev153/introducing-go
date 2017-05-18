package main

import (
	"fmt"
	"time"
)

func pinger(ch chan<- string) {
	for i := 0; ; i++ {
		ch <- "ping"
	}
}

func ponger(ch chan<- string) {
	for i := 0; ; i++ {
		ch <- "pong"
	}
}

func printer(ch <-chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	ch := make(chan string)

	go pinger(ch)
	go ponger(ch)
	go printer(ch)

	var input string
	fmt.Scanln(&input)
}
