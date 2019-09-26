package architecture

import "fmt"

type Person struct {
	First string
}

type Accessor interface {
	Save(n int, p Person)
	Retrieve(n int) Person
}

type PersonService struct {
	a Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{
		a: a,
	}
}

func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("No person with n of %d", n)
	}
	return p, nil
}

func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}

func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}
