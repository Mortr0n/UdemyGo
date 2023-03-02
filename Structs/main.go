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
	age       int
	canDrive  bool
	//? another way is to label the item and then the type instead of just having the type and key be the same
	//? contact contactInfo
	contactInfo
}

func main() {
	jiraiya := person{
		firstName: "Jiraiya",
		lastName:  "Morton",
		age:       15,
		canDrive:  false,
		contactInfo: contactInfo{
			email:   "jmorton@jiraiya.com",
			zipCode: 59802,
		},
	}

	mistyContact := contactInfo{
		email:   "mamorton@morton.com",
		zipCode: 59802,
	}

	misty := person{
		firstName:   "Misty",
		lastName:    "Morton",
		age:         41,
		canDrive:    true,
		contactInfo: mistyContact,
	}
	// fmt.Printf("%+v", jiraiya)

	jiraiya.print()

	misty.print()

	jiraiya.updateName("J", "M")
	mistyPointer := &misty
	fmt.Printf("%+v", mistyPointer)
	fmt.Println()
	fmt.Printf("", mistyPointer) //! craziness.  Accident.  Just kept it because it was kind of cool
	fmt.Println()
	fmt.Printf("%v", mistyPointer)
	fmt.Println(main) //! Getting wild here...
	mistyPointer.updateNameFromPointer("Lookie!", "Whoopie!")

	jiraiya.print()

	misty.print()
	fmt.Println()
	misty.updateName("Look we did it ", "again")
	misty.print()

	//? These are all the ways I declared before adding the contactInfo struct to the mix
	// different ways of declaring and adding values to a struct
	// chris := person{"Chris", "Morton", 42, true}
	// alex := person{firstName: "Alex", lastName: "Anderson", age: 14, canDrive: false}
	// var misty person
	// var jiraiya person

	// misty.firstName = "Misty"
	// misty.canDrive = true

	// jiraiya.firstName = "Jiraiya"
	// jiraiya.age = 15

	// fmt.Println(chris, alex, jiraiya, misty)
	// // %+v will print out all keys and values from the passed in objects
	// fmt.Printf("%+v \n %+v \n %+v \n %+v", alex, misty, chris, jiraiya)
}

// ? feeding the person struct into a receiver function
func (p person) print() {
	fmt.Printf("%+v \n", p)
}

// ? I did the way they showed in the class and the way I originally tried this may have changed
// ? in the last 2 years since the videos were made because both ways work
func (pointerToPerson *person) updateNameFromPointer(newFirstName string, newLastName string) {
	(*pointerToPerson).firstName = newFirstName
	pointerToPerson.lastName = newLastName
}

// ? Go is pass by value.  We need to take a pointer to the instance of the
// ? person stuct in order for us to change the items using this receiver
func (p *person) updateName(firstname string, lastname string) {
	p.firstName = firstname
	p.lastName = lastname
}
