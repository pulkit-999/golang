package main

import (
	"fmt"
	"net/http"
)

func main() {
	// go greeter("Hello")
	// greeter("world")
	websites := []string{
		"https://github.com",
		"https://go.dev",
	}

	for _, web := range websites {
		getStatusCode(web)
	}
}

//	func greeter(s string) {
//		for i := 0; i < 6; i++ {
//			time.Sleep(3 * time.Second)
//			fmt.Println(s)
//		}
//	}
func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d status OK for website:%s\n", res.StatusCode, endpoint)
}
