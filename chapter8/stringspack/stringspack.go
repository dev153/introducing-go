package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test", "st"))
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Repeat("a", 5))
	fmt.Println(strings.Replace("abab", "b", "c", -1))
	fmt.Println(strings.Replace("abab", "b", "c", 1))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	fmt.Println(strings.ToUpper("test"))
	fmt.Println(strings.ToLower("TEST"))
	// fmt.Println(strings.ToTitle("test"))
}
