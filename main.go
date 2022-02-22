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
	// alex := person{"Alex", "Anderson"}

	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)

	// Print the field values of struct "alex"
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 69008,
		},
	}

	jim.updateName("Jimmy")
	jim.print()

}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
