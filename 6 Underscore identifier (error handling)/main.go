package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	result := a / b

	return result, nil

}

func main() {

	fmt.Println("Underscore identifier in GoLang")

	// can also use "_" to ignore the error if we are not interested in printing it
	result, err := divide(10, 0)

	if err != nil {
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Error:", err)
	}

}
