package main

import (
	"fmt"
	"sort"
)

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i].Age, userSlice[j].Age = userSlice[j].Age, userSlice[i].Age
}

func main() {
	users := []User{
		{"Suzuka", 23},
		{"Moa", 21},
		{"Yui", 21},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
