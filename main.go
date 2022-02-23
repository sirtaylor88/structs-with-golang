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

	// Look at this variable and get the memory address of the value
	// this variable is pointing at
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
