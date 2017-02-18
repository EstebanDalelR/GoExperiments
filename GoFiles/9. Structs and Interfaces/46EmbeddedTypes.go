package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is ", p.name)
}

type Android struct {
	Person
	Model string
}

func (a *Android) robotTalk() {
	fmt.Println("0101001", a.name)
}

func main() {
	a := new(Android)
	a.Person.talk()
	a.name = "Roboto"
	a.talk()
}
