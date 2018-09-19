package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	siteList := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
		"http://wikipedia.com",
		"http://reddit.com",
	}

	c := make(chan string)

	for _, link := range siteList {
		go statusCheck(link, c)
	}

	for link := range c {
		go func(l string) {
			time.Sleep(time.Second)
			statusCheck(l, c)
		}(link)
	}

}

func statusCheck(s string, c chan string) {
	_, err := http.Get(s)

	if err == nil {
		fmt.Println(s, " is up!")
	} else {
		fmt.Println(s, " is down!")
	}
	c <- s
}
