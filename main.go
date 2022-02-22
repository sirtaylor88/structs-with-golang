package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{"Alex", "Anderson"}

	// alex := person{firstName: "Alex", lastName: "Anderson"}

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"

	fmt.Println(alex)

	// Print the field values of struct "alex"
	fmt.Printf("%+v", alex)
}
