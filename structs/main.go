package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// p := person{firstName: "Jane", lastName: "Doe"}

	var p person

	p.firstName = "Alex"
	p.lastName = "Anderson"
	p.contact = contactInfo{email: "haha@ahah.com", zipCode: 102030}

	p.print()
	p.update("HAHA")
	p.print()

}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p *person) update(newFN string) {
	p.firstName = newFN
}
