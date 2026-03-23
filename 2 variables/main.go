package main

import "fmt"

func main() {
	fmt.Println("Hello, World")

	var name string = "Mohit Kumhar"
	var version string = "1.0.0"
	fmt.Println(name)
	fmt.Println(version)

	var money int = 67000
	var salary = 67000
	fmt.Println(money)
	fmt.Println(salary)

	var money1 float32 = 67.10
	var salary1 = 67.10
	fmt.Println(money1)
	fmt.Println(salary1)

	var TBoolean bool = true
	var FBoolean = false
	fmt.Println(TBoolean)
	fmt.Println(FBoolean)

	const age = 21
	fmt.Println(age)

	// age = 22
	// fmt.Println(age)

	//  can also be declared like this
	mohit := "Mohit Kumhar"
	fmt.Println(mohit)

	// Public variables start with capital letter and private variables start with small letter
	// public variables can be accessed outside the package but private variables cannot be accessed outside the package

}
