package main

import (
    "fmt"
    "example.com/mymodule/myutil" // ✅ Import using module path
		"example.com/mymodule/mydocx"
)

func main() {
    fmt.Println("Hello, World!")
    myutil.PrintMessage("Hello, world")
		mydocx.Printify("Go is good")
		var name string ="harry"
		fmt.Println(name)
		var version = 76
		fmt.Println(version)
		var cur int = 45
		var currency = 34.7
		fmt.Println("Currency: ", currency)
		fmt.Println("Currency: ", cur)
		var dimension float64 = 798.23
		fmt.Println("Float: ", dimension)
		var person = "Harshit Jain"
		fmt.Println(person)

		champ:="Champion"
		fmt.Println(person)
}
