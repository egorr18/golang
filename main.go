package main

import (
	"fmt"
	"practice/module/interfaces/describer"
	"practice/module/module_1_1/friend"
	"practice/module/module_1_1/me"
	"practice/module/module_1_1/person"
	"practice/module/module_1_1/relative"
)

func main() {
	meUser := me.Me{
		Person: person.Person{
			Name:  "Вовчик",
			Age:   19,
			Hobby: "Клуби",
		},
		Profession: "Розробник",
		Logined:    false,
	}

	friendUser := friend.Friend{
		Person: person.Person{
			Name:  "im_coderrr",
			Age:   20,
			Hobby: "gin",
		},
		FavoriteSong: "daryana, maybe baybe",
	}

	relativeUser := relative.Relative{
		Person: person.Person{
			Name:  "Olexandr",
			Age:   67,
			Hobby: "beer",
		},
		FavoriteMovie: "Тітанік",
	}

	var people []describer.Describer = []describer.Describer{&meUser, &friendUser, &relativeUser}
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

	copyMe := meUser
	copyMe.Logout()
	fmt.Printf("Logined (після Logout копії): %v\n", meUser.Logined)
}
