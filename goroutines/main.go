package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// go greeter("Hello")
	// greeter("world")

	websites := []string{
		"https://github.com",
		"https://go.dev",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
}

//	func greeter(s string) {
//		for i := 0; i < 6; i++ {
//			time.Sleep(3 * time.Second)
//			fmt.Println(s)
//		}
//	}
func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d status OK for website:%s\n", res.StatusCode, endpoint)
}
