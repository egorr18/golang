package person

import "fmt"

type Person struct {
	Name  string
	Age   int
	Hobby string
}

func (p Person) Greet() {
	fmt.Printf("Привіт, мене звати %s, мені %d років\n", p.Name, p.Age)
}

func (p *Person) SetHobby(hobby string) {
	p.Hobby = hobby
}
