package main

import "fmt"

type Animal interface {
	mover() string
}

type Perro struct{}
type Pez struct{}
type Pajaro struct{}

func (*Perro) mover() string {
	return "soy un perro y camino"
}

func (*Pez) mover() string {
	return "soy un pez y nado"
}

func (*Pajaro) mover() string {
	return "soy un pajaro y vuelo "
}
func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}
func main() {
	perror := Perro{}
	moverAnimal(&perror)

	pezr := Pez{}
	moverAnimal(&pezr)

	pajaror := Pajaro{}
	moverAnimal(&pajaror)
}
