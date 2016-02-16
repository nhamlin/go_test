package main

import (
	"fmt"
)

type Person struct {
	Name string // Person has a name
}

type Android1 struct {
	Person Person // Android has a Person object associated with it
	Model  string // Android has a model
}

type Android2 struct {
	Person        // Android IS a person
	Model  string // Android has a model
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func main() {
	a := new(Android1)
	a.Person.Name = "Bob"
	//a.Name is not available
	a.Person.Talk()

	b := new(Android2)
	b.Name = "Molly"
	b.Talk() // The result is Molly
	//b.Name is obviously available
	b.Person.Name = "Delphine"
	//b.Person.Name is also available, and now I'm changing it to be Delphine
	b.Talk() // The result is Delphine, not Molly
}
