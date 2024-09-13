package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	// alex := person{"alex", "anderson"}
	// fmt.Println(alex)
	// ayush := person{firstName: "ayush", lastName: "garg"}
	// fmt.Println(ayush)

	var a person
	a.firstName = "sdvdsvdsvs"
	fmt.Println(a)
	fmt.Printf("%+v",a)

	var jim person
	jim.firstName = "Jim"
	jim.lastName = "Partty"
	jim.contactInfo.email = "dhciudsgc"
	jim.contactInfo.zipCode = 63786284263
	fmt.Println(jim)
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v",p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}