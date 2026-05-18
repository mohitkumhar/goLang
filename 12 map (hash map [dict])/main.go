package main

import "fmt"

func main() {

	//  de
	studentGrades := make(map[string]int)

	// assinging values in the map
	studentGrades["Mohit"] = 34
	studentGrades["Kumhar"] = 12
	studentGrades["Molela"] = 35
	studentGrades["Prajapat"] = 12

	fmt.Println(studentGrades)

	// accessing value from key from map
	fmt.Println("Marks of Mohit is:", studentGrades["Mohit"])
	fmt.Println("Marks of Kumhar is:", studentGrades["Kumhar"])
	fmt.Println("Marks of Molela is:", studentGrades["Molela"])
	fmt.Println("Marks of Prajapat is:", studentGrades["prajapat"])

	// deleting the data from the map
	delete(studentGrades, "Prajapat")
	fmt.Println(studentGrades)

	// checking if key exist or not in map
	prjpt, exists := studentGrades["Prjapat"]
	fmt.Println("Prajapat: ", prjpt)
	fmt.Println("Exists: ", exists)

	// using loop in map
	for key, value := range studentGrades {
		fmt.Printf("Grades of %s is %d\n", key, value)
	}

	// defining values in map when decelaration
	person := map[string]int{
		"Mohit":  12,
		"Kumhar": 24,
	}

	fmt.Println(person)

}
