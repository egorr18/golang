package friend

import (
	"fmt"
	"practice/module/module_1_1/person"
)

type Friend struct {
	person.Person
	FavoriteSong string
}

func (f Friend) Describe() string {
	return fmt.Sprintf("Я — %s. Моє хобі — %s, улюблена пісня — %s", f.Name, f.Hobby, f.FavoriteSong)
}
