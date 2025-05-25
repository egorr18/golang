package person

type Person struct {
	Name  string
	Age   uint8
	Hobby string
}

func NewPerson(name string, age uint8, hobby string) Person {
	return Person{
		Name:  name,
		Age:   age,
		Hobby: hobby,
	}
}
