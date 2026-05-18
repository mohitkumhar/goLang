package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}

		fmt.Println("Numbers is: ", i)
	}

	for {
		fmt.Println("Infinite Loop")
		break
	}

	//  looping on array
	numbers := []int{10, 20, 30, 40, 50}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// looping on string
	name := "Mohit Kumhar"
	for index, char := range name {
		fmt.Printf("Index is: %d, Char is %c\n", index, char)
	}
}
