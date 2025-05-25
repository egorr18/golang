package main

import (
	"KorolEgor/module1/task1/friend"
	"KorolEgor/module1/task1/me"
	"KorolEgor/module1/task1/relative"
	"fmt"
)

type Describer interface {
	Describe() string
}

func main() {
	meUser := me.NewMe("Вовчик", 19, "Клуби", "Розробник")
	friendUser := friend.NewFriend("im_coderrr", 20, "gin", "daryana, maybe baybe")
	relativeUser := relative.NewRelative("Olexandr", 67, "beer", "Тітанік")

	people := [3]Describer{meUser, friendUser, relativeUser}

	for i := range people {
		fmt.Println(people[i].Describe())
	}

	meUser.Greet()
	friendUser.Greet()
	relativeUser.Greet()

	meUser.SetHobby("Клуби")
	friendUser.SetHobby("gin")
	relativeUser.SetHobby("vaib")

	meUser.Login()
	fmt.Printf("Logined (після Login): %v\n", meUser.Logined)

	copyMe := *meUser
	copyMe.Logout()
	fmt.Printf("Logined (після Logout копії): %v\n", meUser.Logined)
}
