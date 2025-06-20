package main

import "fmt"

//Maps in Golang

// In Go, a map is a data structure that provides an unordered collection of key-value pairs, where each key must be unique.

// It is similar to dictionaries or hash maps in other programming languages.

// Maps are used to associate values with keys and allow for efficient retrieval of values based on those keys.
func main() {
	studentGrade := make(map[string]int)

	studentGrade["Prince"] = 34
	studentGrade["Alice"] = 90
	studentGrade["Bob"] = 85
	studentGrade["Charlie"] = 95

	fmt.Println("Marks of Bob:", studentGrade["Bob"])
	studentGrade["Bob"] = 100
	fmt.Println("new Marks of Bob: ", studentGrade["Bob"])

	delete(studentGrade, "Bob")
	fmt.Println("Masrks of Bob: ", studentGrade["Bob"])
	//cheacking if a key exits
	grades, exists := studentGrade["David"]
	fmt.Println("Grades of David : ", grades)
	fmt.Println("David exits: ", exists)

	Grades, Exists := studentGrade["Prince"]
	fmt.Println("Grades of David : ", Grades)
	fmt.Println("Prince exits: ", Exists)

	for index, value := range studentGrade {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	person := map[string]int{
		"alice":   90,
		"bob":     85,
		"Charlie": 95,
	}
	for index, value := range person {
		fmt.Printf("------Key is %s and marks is %d\n", index, value)
	}
}
