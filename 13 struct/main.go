package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	City string
	Pin  int
}

type Employee struct {
	Person_details Person
	Person_contact Contact
	Person_address Address
}

func main() {
	fmt.Println("Structure in GoLang")

	var mohit Person
	mohit.FirstName = "Mohit"
	mohit.LastName = "Kumhar"
	mohit.Age = 21

	fmt.Println("Mohit Person: ", mohit)

	person1 := Person{
		FirstName: "Umesh",
		LastName:  "Kumhar",
		Age:       21,
	}

	fmt.Println(person1)

	// defining using `new` keyword
	var person2 = new(Person)
	person2.FirstName = "Prajapat"
	person2.LastName = "Molela"
	person2.Age = 21

	fmt.Println(person2)
	fmt.Println(person2.FirstName)
	fmt.Println()
	fmt.Println()

	employee1 := new(Employee)

	employee1.Person_details.FirstName = "Mohit"
	employee1.Person_details.LastName = "Kumhar"
	employee1.Person_details.Age = 21

	employee1.Person_contact.Email = "mohitmolela@gmail.com"
	employee1.Person_contact.Phone = "+1234567890"

	employee1.Person_address.City = "Udaipur"
	employee1.Person_address.Pin = 123456

	fmt.Println(employee1)

}
