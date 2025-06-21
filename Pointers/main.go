package main

import "fmt"

//Pointers in Golang
// In Go, a pointer is a variable that stores the memory address of another variable.

// Pointers are used to indirectly refer to the value stored in a variable, rather than the value itself.

// They provide a way to work with the memory directly, which can be useful for various programming tasks, including efficient memory management and sharing data between functions.

// To declare a pointer, you use the * (asterisk) symbol followed by the type of the variable it will point to. You can then initialize the pointer using the address-of (&) operator.
func modifyValueByReference(num *int) {
	*num = *num * 20
}
func main() {
	// var num int
	// num = 2
	num := 2
	// var ptr *int
	// ptr = &num
	ptr := &num

	fmt.Println("Num has value: ", num)
	fmt.Println("pointer contains: ", ptr)
	fmt.Println("Data contains through Pointer : ", *ptr)

	var pointer *int
	if pointer = nil {
		fmt.Println("Pointer is not assigned")
	}

	value := 10
	modifyValueByReference(&value)
	fmt.Println("Value contains : ", value)
}
