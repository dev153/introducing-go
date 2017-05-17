package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

//=============================================================================

type ByAge []Person

func (ps ByAge) Len() int {
	return len(ps)
}

func (ps ByAge) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

func (ps ByAge) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

//=============================================================================

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		Person{Name: "Jill", Age: 11},
		Person{Name: "Avery", Age: 12},
		Person{Name: "Jack", Age: 10},
	}

	fmt.Println(kids)
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
	sort.Sort(ByName(kids))
	fmt.Println(kids)
}
