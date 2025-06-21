package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three two two three two one two two"
	count := strings.Count(str, "two")
	fmt.Println("count: ", count)

	str = " hello, go"
	trimmed := strings.TrimSpace(str)
	fmt.Println("Trimmed: ", trimmed)

	str1 := "Harshit"
	str2 := "Jain"
	result := strings.Join([]string{str1, str2}, " ")
	fmt.Println("result: ", result)
}
