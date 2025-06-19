package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    fmt.Println("Hey, what's your name?")
    
    //var name string
    //fmt.Scan(&name)
    //fmt.Println("Hello, Mr.", name) // This is basic input without spaces

    // Now using bufio to read full line (e.g., full name with spaces)
    fmt.Println("Enter your full name:")
    reader := bufio.NewReader(os.Stdin)
    name, _ := reader.ReadString('\n') // ✅ use '=' not ':='
    fmt.Println("Hello, Mr.", name)
}
