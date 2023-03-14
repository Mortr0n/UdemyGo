package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		go checkLink(link)
	}

}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " errored and might be down!")
		return
	}

	fmt.Println(link + " appears to be working!")
}
