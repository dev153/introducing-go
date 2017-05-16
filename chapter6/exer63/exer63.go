package main

import (
	"fmt"
)

func findGreatest(args ...int) (int, string) {
	err := ""
	result := 0
	lenofargs := len(args)
	if lenofargs == 0 {
		err = "Number of arguments is zero. Can't find the greatest value."
		return 0, err
	}

	result = args[0]

	if lenofargs != 1 {
		for i := 1; i < lenofargs; i++ {
			if args[i] > result {
				result = args[i]
			}
		}
	}

	return result, err
}

func main() {

	printResult := func(result int, err string) {
		if err == "" {
			fmt.Println(result)
		} else {
			fmt.Println(err)
		}
	}

	result, err := findGreatest()
	printResult(result, err)
	result, err = findGreatest(1, 2, 3, 4)
	printResult(result, err)
	result, err = findGreatest(-4, -3, -2, -1)
	printResult(result, err)

}
