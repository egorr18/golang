package me

import (
	"fmt"
	"practice/module/person"
)

type Me struct {
	person.Person
	Profession string
	Logined    bool
}

func NewMe(name string, age int, hobby, profession string) *Me {
	return &Me{
		Person: person.Person{
			Name:  name,
			Age:   age,
			Hobby: hobby,
		},
		Profession: profession,
	}
}

func (m Me) Describe() string {
	return fmt.Sprintf("Я %s, %s, моє хобі — %s", m.Name, m.Profession, m.Hobby)
}

func (m *Me) Login() {
	m.Logined = true
}

func (m *Me) Logout() {
	m.Logined = false
}
