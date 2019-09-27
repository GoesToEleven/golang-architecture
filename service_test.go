package architecture

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

type Db map[int]Person

func (m Db) Save(n int, p Person) {
	m[n] = p
}

func (m Db) Retrieve(n int) Person {
	return m[n]
}

func TestPut(t *testing.T) {
	ctl := gomock.NewController(t)
	acc := NewMockAccessor(ctl)

	p := Person{
		First: "James",
	}

	acc.EXPECT().Save(1, p).MinTimes(1).MaxTimes(1)

	Put(acc, 1, p)

	ctl.Finish()
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
