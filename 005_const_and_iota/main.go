package main

import "fmt"

type Role int

const (
	Admin Role = iota
	Moderator
	User
)

func (r Role) String() string {
	switch r {
	case Admin:
		return "Admin"
	case Moderator:
		return "Moderator"
	case User:
		return "User"
	default:
		return "Unknown"
	}
}

func main() {
	var r Role = Moderator
	fmt.Println(r)           // Moderator
	fmt.Println(Admin, User) // 0 2 (nilai underlying int)
}

