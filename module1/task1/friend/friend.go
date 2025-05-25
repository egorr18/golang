package friend

import (
	"fmt"

	"KorolEgor/module1/task1/person"
)

type Friend struct {
	Person person.Person
	Phone  string
}

func NewFriend(name string, age uint8, hobby string, phone string) *Friend {
	return &Friend{
		Person: person.NewPerson(name, age, hobby),
		Phone:  phone,
	}
}

func (f *Friend) Describe() string {
	return fmt.Sprintf("%s, %d років, hobby: %s, phone: %s", f.Person.Name, f.Person.Age, f.Person.Hobby, f.Phone)
}

func (f *Friend) Greet() {
	fmt.Println("Привіт від друга:", f.Person.Name)
}

func (f *Friend) SetHobby(hobby string) {
	f.Person.Hobby = hobby
}
