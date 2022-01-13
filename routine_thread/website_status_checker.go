package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	/*
		http://google.com is up!
		http://facebook.com is up!
		http://stackoverflow.com is up!
		http://golang.org is up!
		http://amazon.com is up!
	*/
	for i, link := range links {
		fmt.Println(i)
		checkLink(link)
	}
}

func checkLink(link string) {
	_, error := http.Get(link) // 不在乎response所以response改为下划线_，只保留error
	if error != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}

