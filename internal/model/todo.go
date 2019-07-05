package model

type Ownable interface {
	Owner() *User
}

type Todo struct {
	ID      int
	Text    string
	Done    bool
	MyOwner *User
}

func (t Todo) Owner() *User {
	return t.MyOwner
}

type User struct {
	ID   int
	Name string
}
