package main

import (
	"fmt"
	"strings"
)

var allNames []string

type User struct {
	names []int
}

func NewUser(name string) *User {
	getOrAdd := func(s string) int {
		for i := range allNames {
			if allNames[i] == s {
				return i
			}
		}
		allNames = append(allNames, s)
		return int(len(allNames) - 1)
	}

	result := User{}
	parts := strings.Split(name, " ")
	for _, v := range parts {
		result.names = append(result.names, getOrAdd(v))
	}
	return &result
}

func (u *User) Fullname() string {
	var parts []string
	for _, v := range u.names {
		parts = append(parts, allNames[v])
	}
	return strings.Join(parts, " ")
}

func main() {
	u1 := NewUser("Clyne Fuentes")
	u2 := NewUser("Clippe Cruz")
	u3 := NewUser("Claire Cruz")
	fmt.Println(u1.Fullname())
	fmt.Println(u2.Fullname())
	fmt.Println(u3.Fullname())
}