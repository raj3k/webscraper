package main

import (
	"fmt"
	"os"
	"webscraper/internal/client"
)

func main() {
	url := "https://example.com/"
	response, err := client.Get(url)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	fmt.Println(response)
}
