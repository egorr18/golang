package relative

import (
	"fmt"
	"practice/module/module_1_1/person"
)

type Relative struct {
	person.Person
	FavoriteMovie string
}

func (r Relative) Describe() string {
	return fmt.Sprintf("Мене звати %s, моє хобі — %s, улюблений фільм — %s", r.Name, r.Hobby, r.FavoriteMovie)
}
