package main

import "fmt"

func main() {
	fmt.Println("############## map 写法1 ##############")
	colors := map[string]int{
		"red":   49,
		"green": 2139,
	}

	fmt.Println(len(colors)) // 2
	/*
		############## delete for map ##############
	*/
	delete(colors, "green")
	fmt.Println(len(colors)) // 1
}
