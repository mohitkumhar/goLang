// An array is a collection of elements of the same type, stored in contiguous memory locations.
// In this program, we declare an array of strings to store names and an array of integers to store numbers.
// We also print the contents of the arrays and their lengths.

package main

import "fmt"

func main() {
	fmt.Println("Array In GoLang")

	var name [5]string

	name[0] = "Mohit"
	name[1] = "Kumhar"
	name[2] = "Molela"
	name[4] = "Prajapat"

	fmt.Println(name)
	fmt.Printf("Name Array is: %q\n", name)

	//  can also be declared using
	var numbers = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Numbers is: ", numbers)
	fmt.Println("Length of number array is: ", len(numbers))
	fmt.Println("value of index 2 is: ", numbers[2])
}
