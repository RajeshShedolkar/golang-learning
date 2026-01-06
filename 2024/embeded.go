package main

import "fmt"

// A struct's fields usually represent the has-a relationship. 
// For example a Circle has a radius.
//

type Person struct {
	Name string
}

type Android struct {
	Person // Person Person, P Person
	Model  string
}

func main() {
	a := new(Android)
	a.Person.Name = "Rajesh"
	a.Person.talk()
	a.talk()
}

func (p *Person) talk() {
	fmt.Println("I am Talking. and My name is", p.Name)
}
