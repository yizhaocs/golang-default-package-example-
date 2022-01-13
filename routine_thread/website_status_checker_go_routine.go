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

	channel := make(chan string) // channel 用string作为data type做message 交流

	/*
		http://google.com is up!
		http://facebook.com is up!
		http://stackoverflow.com is up!
		http://golang.org is up!
		http://amazon.com is up!
	*/
	for i, link := range links {
		fmt.Println(i)
		go checkLink_with_channel(link, channel)
	}

	fmt.Println(<-channel)
}

func checkLink_with_channel(link string, channel chan string) {
	_, error := http.Get(link) // 不在乎response所以response改为下划线_，只保留error
	if error != nil {
		fmt.Println(link, "might be down!")
		channel <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	channel <- "Yep its up"
}
