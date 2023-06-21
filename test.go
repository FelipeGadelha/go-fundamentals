package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u *User) convert(user User) {
	u.name = user.name
	u.age = user.age
}

func main() {
	var user User = User{
		name: "felipe",
		age:  21,
	}

	user2 := User{}
	user2.convert(user)

	fmt.Println(user2)
}
