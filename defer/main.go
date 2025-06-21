package main

import "fmt"

func main() {
	fmt.Println("Starting of the program")
	defer fmt.Println("Middle1 of the program")
	defer fmt.Println("Middle2 of the program")
	fmt.Println("End of the program")
}
