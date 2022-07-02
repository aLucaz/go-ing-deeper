package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func main() {
	// var alex person
	alex := person{
		firstName: "alex",
		lastName:  "martinez",
	}

	fmt.Printf("%+v\n", alex)
	fmt.Println("=============================")
	alex.contact.email = "alex.ocm"
	alex.contact.zip = 123

	fmt.Printf("%+v\n", alex)

	alexPointer := &alex
	alexPointer.updateName("arturo")
	alex.print()

	// *person -> pointer of type person
	// &alex -> memory address of alex
	// *alex -> value that &alex points

	// a shortcut
	alex.updateName("lucas")
	alex.print()
}
