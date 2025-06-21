package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Learning web service")

	// Perform HTTP GET request
	res, err := http.Get("http://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response:", err)
		return
	}
	defer res.Body.Close()

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print response as string
	fmt.Println("Response:", string(data))
}
