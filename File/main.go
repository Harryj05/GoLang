package main

import (
	"fmt"
	"os"
)

func main() {
	/*
			file, err := os.Create("example.txt")
			if err != nil {
				fmt.Println("Error while creating file: ", err)
				return
			}
			defer file.Close()

			content := "hello world by harshit"
			byte, error := io.WriteString(file, content+"\n")
			fmt.Println("byte written: ", byte)
			if error != nil {
				fmt.Println("Error while writing file: ", error)
			}
			fmt.Println("successfully created file")

		file, err := os.Open("example.txt")
		if err != nil {
			fmt.Println("Error while opening file: ", err)
			return
		}
		defer file.Close()
		//Create a buffer to read the file content
		buffer := make([]byte, 1024)
		//Read the file content into the buffer
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error while reading file", err)
				return
			}
			//Process the read content
			fmt.Println(string(buffer[:n]))
		}*/
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file ")
		return
	}
	fmt.Println(string(content))
}
