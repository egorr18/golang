package relative

import (
	"fmt"
	"practice/module/person"
)

type Relative struct {
	person.Person
	FavoriteMovie string
}

func NewRelative(name string, age int, hobby, movie string) *Relative {
	return &Relative{
		Person: person.Person{
			Name:  name,
			Age:   age,
			Hobby: hobby,
		},
		FavoriteMovie: movie,
	}
}

func (r Relative) Describe() string {
	return fmt.Sprintf("Мене звати %s, моє хобі — %s, улюблений фільм — %s", r.Name, r.Hobby, r.FavoriteMovie)
}
