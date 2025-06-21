package main

//net/url
import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")

	myURL := "https://chatgpt.com/c/68559fe3-a570-8001-8db0-a26a34a5b5f4"

	fmt.Printf("Type of URL %T\n", myURL)

	parseURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Can't parse URL")
		return
	}
	fmt.Printf("Type of URL: %T\n", parseURL)

	fmt.Println("Scheme: ", parseURL.Scheme)
	fmt.Println("Host: ", parseURL.Host)
	fmt.Println("Port: ", parseURL.Path)
	fmt.Println("Scheme: ", parseURL.RawQuery)

	//odifying URL components
	parseURL.Path = "/newPath"
	parseURL.RawQuery = "username=imageharry"

	newUrl := parseURL.String()
	fmt.Println("new url ", newUrl)
}
