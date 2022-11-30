package main

import "fmt"

func main() {
	p := new(Person)

	var if1 If1 = p
	b := if1.Less(1)
	if1.Add(1)
	fmt.Printf("%v, %v", b, if1)
}

type Person struct {
	age int
}

func (p Person) Less(i int) bool {
	return p.age < i
}
func (p *Person) Add(i int) {
	p.age += i
}

type If1 interface {
	Less(i int) bool
	Add(i int)
}
