package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for item := x.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value.(int))
	}
}
