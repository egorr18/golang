package me

import (
	"fmt"

	"KorolEgor/module1/task1/person"
)

type Me struct {
	Person  person.Person
	Job     string
	Logined bool
}

func NewMe(name string, age uint8, hobby, job string) *Me {
	return &Me{
		Person:  person.NewPerson(name, age, hobby),
		Job:     job,
		Logined: false,
	}
}

func (m *Me) Describe() string {
	return fmt.Sprintf("%s, %d років, hobby: %s, job: %s", m.Person.Name, m.Person.Age, m.Person.Hobby, m.Job)
}

func (m *Me) Greet() {
	fmt.Println("Привіт від мене:", m.Person.Name)
}

func (m *Me) SetHobby(hobby string) {
	m.Person.Hobby = hobby
}

func (m *Me) Login() {
	m.Logined = true
}

func (m *Me) Logout() {
	m.Logined = false
}
