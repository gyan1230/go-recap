package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"http://www.google.com",
		"http://34.80.248.226/login",
		"http://www.yahoo.com",
		"http://www.facebook.com",
		"http://www.golang.org",
	}

	// for _, url := range urls {
	// 	status, err := check(url)
	// 	if err != nil {
	// 		fmt.Printf("Error: %s\n", err)
	// 	}
	// 	fmt.Printf("%s Status code : %d\n", url, status)
	// }

	resp := make(chan string)
	for _, url := range urls {
		go checker(url, resp)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s", <-resp)
	}
}

func check(url string) (status int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

func checker(url string, resp chan string) {
	res, err := check(url)
	if err == nil {
		resp <- fmt.Sprintf("%s Status code : %d\n", url, res)
	}
}
