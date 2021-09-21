package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	// dwight := person{"Dwight", "Schrute"}
	// dwight := person{
	// 	firstName: "Dwight",
	// 	lastName: "Schrute",
	// }
	// var dwight person
	// dwight.firstName = "Dwight"
	// dwight.lastName = "Schrute"
	// fmt.Println(dwight)

	dwight := person{
		firstName: "Dwight",
		lastName:  "Schrute",
		contactInfo: contactInfo{
			email:   "star_salesman@sabre.com",
			zipCode: 80000,
		},
	}
	// fmt.Printf("%+v", dwight)
	// dwight.print()

	// dwightPointer := &dwight
	// dwightPointer.updateName("Beets")

	dwight.updateName("Beets")
	dwight.print()
}

func (pointer *person) updateName(newFirstName string) {
	(*pointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
