package architecture

import (
	"fmt"
	"testing"
)

type Db map[int]Person

func (m Db) Save(n int, p Person) {
	m[n] = p
}

func (m Db) Retrieve(n int) Person {
	return m[n]
}

func TestPut(t *testing.T) {
	mdb := Db{}
	p := Person{
		First: "James",
	}

	Put(mdb, 1, p)

	got := mdb.Retrieve(1)
	if got != p {
		t.Fatalf("Want %v, got %v", p, got)
	}
}

func ExamplePut() {
	mdb := Db{}
	p := Person{
		First: "James",
	}

	Put(mdb, 1, p)
	got := mdb.Retrieve(1)
	fmt.Println(got)
	// Output: {James}
}
