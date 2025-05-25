package relative

import (
	"KorolEgor/module1/task1/person"
	"fmt"
)

type Relative struct {
	Person person.Person
	Note   string
}

func NewRelative(name string, age uint8, hobby, note string) *Relative {
	return &Relative{
		Person: person.NewPerson(name, age, hobby),
		Note:   note,
	}
}

func (r *Relative) Describe() string {
	return fmt.Sprintf("%s, %d років, hobby: %s, note: %s", r.Person.Name, r.Person.Age, r.Person.Hobby, r.Note)
}

func (r *Relative) Greet() {
	fmt.Println("Привіт від родича:", r.Person.Name)
}

func (r *Relative) SetHobby(hobby string) {
	r.Person.Hobby = hobby
}
