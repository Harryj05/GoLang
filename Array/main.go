package main

import "fmt"

func main() {
	fmt.Println("we are learning Array in Golang")

	// var name [5]string

	// //0 1 2 3 4
	// name[0] = "Harshit"
	// name[2] = "Sanagam"
	// name[3] = "Hritvik"

	// fmt.Println("Names of Person is:", name)

	// var numbers = [5]int{1, 2, 3, 4, 5}
	// fmt.Println("Number is :", numbers)

	// fmt.Println("Length of numbers array is :", len(numbers))

	// fmt.Println("value of name at 2 index is:", name[2])
	var price [5]int
	fmt.Println("Price is:", price)
	fmt.Printf("Price array is %q\n", price)
}
