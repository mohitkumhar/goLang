//  print and println are used to print the output to the console.
// The difference between them is that println adds a new line after printing the output, while print does not.
// In the below code, we are using both print and println to demonstrate their usage.

package main

import "fmt"

func main() {
	age := 21
	name := "Mohit Kumhar"
	height := 5.14560

	fmt.Println("Hello, World")
	fmt.Println(age)
	fmt.Print(name)
	fmt.Println(height)

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Name is %s\n", name)
	fmt.Printf("Height is %.3f\n", height)

	fmt.Printf("Height is %T\n", height) // for printing the type of datatype
}
