package architecture

type User struct {
	First string
}

// Accessor is how to access saved users
type Accessor interface {
	Save(n int, u User)
	Retrieve(n int) User
}

func Put(a Accessor, n int, u User) {
	a.Save(n, u)
}

func Get(a Accessor, n int) User {
	return a.Retrieve(n)
}
