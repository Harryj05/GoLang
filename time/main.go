package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("Current time : ", currentTime)

	fmt.Printf("Typ of currentTime %T\n", currentTime)
	formatted := currentTime.Format("02-01-2006, 14:04:05")
	fmt.Println("formatted time: ", formatted)

}
