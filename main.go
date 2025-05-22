package main

import (
	"fmt"

	"practice/module/describer"
	"practice/module/friend"
	"practice/module/me"
	"practice/module/relative"
)

func main() {
	meUser := me.NewMe("Вовчик", 19, "Клуби", "Розробник")
	friendUser := friend.NewFriend("im_coderrr", 20, "gin", "daryana, maybe baybe")
	relativeUser := relative.NewRelative("Olexandr", 67, "beer", "Тітанік")

	people := []describer.Describer{meUser, friendUser, relativeUser}

	for _, p := range people {
		fmt.Println(p.Describe())
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
