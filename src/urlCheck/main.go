package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResults struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

// functional programing
func main() {
	// map이 초기화가 안되서 에러
	// var requestResultss map[string]string

	// make() 함수로 초기화
	var results = make(map[string]string)
	c := make(chan requestResults)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		// "https://academy.nomadcoders.co/",
	}

	// for _, url := range urls {
	// 	requestResults := "OK"
	// 	err := hitURL(url)
	// 	if err != nil {
	// 		requestResults = "Failed"
	// 	}
	// 	requestResultss[url] = requestResults
	// }
	// for url, requestResults := range requestResultss {
	// 	fmt.Println(url, requestResults)
	// }

	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		// fmt.Println(<-c)
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}

}

// func hitURL(url string) error {
// 	// 이런 방식은 동기적으로 한번 그다음 한번 실행한다.
// 	fmt.Println("Checking: ", url)
// 	resp, err := http.Get(url)
// 	if err != nil || resp.StatusCode >= 400 {
// 		fmt.Println(err)
// 		return errRequestFailed

// 	}
// 	return nil
// }

func hitURL(url string, c chan<- requestResults) {
	// fmt.Println("Checking: ", url)
	// this channel is send only
	// fmt.Println(<-c)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResults{url: url, status: status}
}
