package friend

import (
	"fmt"
	"practice/module/person"
)

type Friend struct {
	person.Person
	FavoritePerson string
}

func NewFriend(name string, age int, hobby, favoritePerson string) *Friend {
	return &Friend{
		Person: person.Person{
			Name:  name,
			Age:   age,
			Hobby: hobby,
		},
		FavoritePerson: favoritePerson,
	}
}

func (f Friend) Describe() string {
	return fmt.Sprintf("Я %s, люблю %s, моє хобі — %s", f.Name, f.FavoritePerson, f.Hobby)
}
