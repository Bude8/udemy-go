package main

import "fmt"

// Defining structs
// Embedding struct
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo can embed struct this way too
}

func main() {
	// Declaring structs
	// alex := person{"Alex", "Anderson"} // not recommended as order may change
	// alex := person{firstName: "Alex", lastName: "Anderson"} // recommended
	// var alex person // "zero value" given if unassigned, e.g. "", 0, false
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{firstName: "Jim",
		lastName: "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	// jim := person{firstName: "Jim",
	// 	lastName: "Party",
	// 	contactInfo: contactInfo{
	// 		email:   "jim@gmail.com",
	// 		zipCode: 12345,
	// 	},
	// }

	// Pointers!
	// & - give me the memory address of the value this variable is pointing at
	// * - give me the value this memory address is pointing at

	// SHORTCUT
	// If we have a variable of type person, we can have a receive of pointer to person, that's ok
	// jim is type of person
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	// Pass by value
	// when this is called, its value is copied to another memory address
	// then the receiver p points to the copy
	jim.print()

	// GOTCHAS!!
	// Slices do not need pointers, structs do!
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

	// Arrays
	// - Primitive data structure
	// - Can't be resized
	// - Rarely used directly

	// Slices
	// - Can grow and shrink
	// - Used 99% of the time for lists of elements

	// Creaeting a slice of string in go creates two different data structures - slice and array
	// Slice
	// - pointer to array head,
	// - capacity - how many can be in the array,
	// - length - how long the array currently is
	// The slice and array are stored at different addresses but the slice points
	// to the memory address of the array
	// Go still passes by value - but does so for the slice. However, the slice is pointing at the array still.
	// Slices, maps, channels, pointers, functions - reference types - don't worry about pointers with these
	// int, float, string, bool, structs - value types - use pointers to change these things in a function
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// * in front of a type is a TYPE DESCRIPTION
// it means we are working with a pointer to a person
// * in front of anything else - it is an OPERATOR
// it means we want to manipulate the value the pointer is referencing
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Turn address into value with *address
// Turn value into address with &value

func updateSlice(s []string) {
	s[0] = "Bye"
}
