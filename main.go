package main

import (
	"fmt"
)

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

func main() {
	alex := User{
		nombre: "Alex",
		email:  "alex@gmail.com",
		activo: true,
	}
	ivan := User{
		nombre: "Ivan",
		email:  "ivan@gmail.com",
		activo: true,
	}
	alexStud := Student{
		user:   alex,
		codigo: "001",
	}
	fmt.Println(ivan)
	fmt.Println(alexStud)
}
