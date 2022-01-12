package main

import (
	"fmt"
)

func main() {
	{
		fmt.Println("############## map 写法1 ##############")
		colors := map[string]int{
			"red":   49,
			"green": 2139,
		}

		fmt.Println(colors) // map[green:2139 red:49]
		/*
			delete for map
		*/
		delete(colors, "green")
		fmt.Println(colors) // map[red:49]
		/*
			append for map
		*/
		colors["black"] = 23321
		fmt.Println(colors) // map[black:23321 red:49]
		/*
			Iterating Over Maps
			k: red v: 49
			k: black v: 23321
		*/
		printMap(colors)
	}

	{
		fmt.Println("############## map 写法2 ##############")
		var colors2 map[string]int
		fmt.Println(colors2) // map[]
		colors2 = make(map[string]int)
		colors2["pink"] = 213
		fmt.Println(colors2) // map[pink:213]
	}
	{
		fmt.Println("############## map 写法3 ##############")
		colors3 := make(map[string]int)
		colors3["white"] = 213
		fmt.Println(colors3) // map[white:213]
	}
}

func printMap(inputMap map[string]int) {
	for k, v := range inputMap {
		fmt.Println("k:", k, "v:", v)
	}
}
