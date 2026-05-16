package main

import "fmt"

func simpleFunction() {
	fmt.Println("This is a simple function")
}

func add(a int, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Println("Functions in GoLang")

	simpleFunction()

	ans := add(5, 10)
	fmt.Println("The sum of 5 and 10 is:", ans)

}
