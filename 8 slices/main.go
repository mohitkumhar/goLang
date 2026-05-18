package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Original slice:", numbers)
	fmt.Printf("Number has datatype of: %T\n", numbers)
	fmt.Println("Length: ", len(numbers))

	numbers = append(numbers, 6, 7, 8)
	fmt.Println("Original slice:", numbers)
	fmt.Println("Length: ", len(numbers))
	fmt.Println("Capacity: ", cap(numbers))

	num := make([]int, 5, 10)
	fmt.Println("New slice:", num)
	fmt.Println("Length: ", len(num))
	fmt.Println("Capacity: ", cap(num))

	num = append(num, 1, 2, 3, 4, 5)
	fmt.Println("New slice:", num)
	fmt.Println("Length: ", len(num))
	fmt.Println("Capacity: ", cap(num))

	num = append(num, 6)
	fmt.Println("New slice:", num)
	fmt.Println("Length: ", len(num))
	fmt.Println("Capacity: ", cap(num))
}
