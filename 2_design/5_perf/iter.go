package main

import "fmt"

type User struct {
	Login string
	// TODO: More fields
}

type UserIter interface {
	Next(*User) bool
}

type Query struct {
	n int
}

func (q *Query) Next(u *User) bool {
	if q.n == 5 {
		return false
	}

	q.n++
	u.Login = fmt.Sprintf("user-%d", q.n)

	return true
}

func PrintUsers(ui UserIter) {
	var u User
	for ui.Next(&u) {
		fmt.Println(u)
	}
}

func main() {
	var q Query
	PrintUsers(&q)
}
