package main

import "fmt"

type Person struct {
	Name string
}

type Android struct {
	Person
	Model string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	personA := Person{"Nick"}
	personA.Talk()
	androidA := Android{Model: "R2-D2"}
	fmt.Println(androidA)
	androidA.Person.Talk()
	androidB := new(Android)
	androidB.Talk()
}
