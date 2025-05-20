package me

import (
	"fmt"
	"practice/module/module_1_1/person"
)

type Me struct {
	person.Person
	Profession string
	Logined    bool
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
