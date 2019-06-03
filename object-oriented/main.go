package main

import "fmt"

type Animal struct {
	Name   string
	canFly bool
}

type Person struct {
	name string
	age  int
}

func (p Person) canVote() int {
	return p.age
}

type Pigeon struct {
	Name          string
	featherLength int
}

func (p *Pigeon) GetFeatherLength() int {
	return p.featherLength
}
func (p *Pigeon) SetFeatherLength(length int) {
	p.featherLength = length
}

func main() {
	anAnimal := Animal{Name: "Lion", canFly: false}
	fmt.Println(anAnimal.Name)
	aLionPtr := &anAnimal
	fmt.Println(aLionPtr.canFly)

	p := Pigeon{"Tweety", 10}
}
