package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p *person) speak() {
	fmt.Println("from a person", p.first)
}

func (sa *secretAgent) speak() {
	fmt.Println("this is not my real name b/c I'm a secret agent", sa.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Jenny",
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	saySomething(&p1)
	saySomething(&sa1)
}
