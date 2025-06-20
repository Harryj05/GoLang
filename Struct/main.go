package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var prince Person
	fmt.Println("Prince Person : ", prince)
	prince.FirstName = "Prince"
	prince.LastName = "Agarwal"
	prince.Age = 24
	fmt.Println("Prince Person : ", prince)

	person1 := Person{
		FirstName: "Akash",
		LastName:  "Jain",
		Age:       23,
	}
	fmt.Println("Person 1: ", person1)

	//new keyword
	var person2 = new(Person)
	person2.FirstName = "Simaran"
	person2.LastName = "burgers"
	person2.Age = 23

	fmt.Println("Person 2 : ", person2)
}
