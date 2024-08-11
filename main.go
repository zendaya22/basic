package main

import (
	"fmt"

)

type profile struct {
	name string
	age int
	address string
	sosmed string
	contact string
}

type account struct{
	*profile
	username string
	password int
}

func getUser(p profile){
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.address)
	fmt.Println(p.contact)
}



func main(){

	user := profile{
		name: "dimas",
		age: 20,
		address: "bumijawa",
		contact: "081717316596",
	}

	acc := account{
		profile: &user,
		username: "dimas anjay mabar",
		password: 11212,
	}

	getUser(user)
	fmt.Println(*acc.profile)
}