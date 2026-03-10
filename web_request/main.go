package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://prettier.io/"

func main() {
	fmt.Println("hI")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	fulldata := string(databytes)
	fmt.Println(fulldata)
}
