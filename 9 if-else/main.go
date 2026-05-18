package main

import "fmt"

func main() {
	age := 18

	if age <= 18 {
		fmt.Println("Age is less than 18")
	} else {
		fmt.Println("Age is Greater then 18")
	}

	new_age := 18

	if new_age < 18 {
		fmt.Println("Your Age is Less then 18")
	} else if new_age == 18 {
		fmt.Println("Your are of 18 Year age")
	} else {
		fmt.Println("You are greater then 18")
	}

	time := 10

	if time < 12 && time > 6 {
		fmt.Println("Good Morning")
	} else {
		fmt.Println("Good Evening")
	}

}
