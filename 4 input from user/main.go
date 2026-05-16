package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string

	// fmt.Println("Enter your name:")
	// fmt.Scan(&name)

	// // it will not take input after the blank space
	// fmt.Printf("Hello, %s\n", name)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your full name:")
	fullName, _ := reader.ReadString('\n')
	fmt.Printf("\nHello, %s\n", fullName)

}
