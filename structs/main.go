package main

import "fmt"

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
	vlad := person{
		firstName: "Vlad",
		lastName:  "Manyechukvyk",
		contact: contactInfo{
			email:   "vm@gmail.com",
			zipCode: 121321,
		},
	}
	vladPointer := &vlad
	vladPointer.updateName("Vladium")
	vladPointer.print()
}

func (personPointer *person) updateName(newFirstName string) {
	(*personPointer).firstName = newFirstName
}

func (personPointer person) print() {
	fmt.Printf("%+v", personPointer)
}
