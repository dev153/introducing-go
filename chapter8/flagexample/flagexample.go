package main

import (
	"flag"
	"fmt"
)

func main() {
	maxp := flag.Int("max", 10, "The max value")
	guiMode := flag.Bool("gui", false, "Start application in GUI")
	fmt.Println(*maxp)
	fmt.Println(*guiMode)
}
