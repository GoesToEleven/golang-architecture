package main

import "fmt"

type person struct {
	first string
}

func (p person) String() string {
	return fmt.Sprintf("My name is %s", p.first)
}

func main() {
	p1 := person{
		first: "Jenny",
	}

	p2 := person{
		first: "James",
	}

	m := map[person]string{}

	m[p1] = "loves james bond"
	m[p2] = "loves james bond also"

	fmt.Println(m[p1])
	fmt.Println(m[p2])
	fmt.Println(p1==p2)
	fmt.Println(p1==p1)
}
