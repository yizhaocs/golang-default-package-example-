package main

import "fmt"

func main() {
	{
		colors := map[string]int{
			"red":   49,
			"green": 2139,
		}

		fmt.Println(colors) // map[green:2139 red:49]
		/*
			############## delete for map ##############
		*/
		delete(colors, "green")
		fmt.Println(colors) // map[red:49]
	}
}
