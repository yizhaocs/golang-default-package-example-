package main

import "fmt"

func main() {
	{
		{
			fmt.Println("############## string 写法1 ##############")
			s2 := "abcd"
			fmt.Println("string:", s2) // string: abcd
		}
		{
			fmt.Println("############## string 写法2 ##############")
			var s1 string
			s1 = "s1"
			fmt.Println("string:", s1) // string: s1
		}
	}
	{
		mySlice := []string{"Hi", "There", "how", "are", "you?"}
		fmt.Println("Slice:", mySlice) // Slice: [Hi There how are you?]
	}
	{
		var i1 int
		i1 = 3
		fmt.Println("int:", i1) // int: 3
	}
	{
		var f1 float32
		f1 = 3.4
		fmt.Println("float32:", f1) // float32: 3.4
	}
	{
		var f1 float64
		f1 = 4.3
		fmt.Println("float64:", f1) // float64: 4.3
	}
	{
		var b1 bool
		b1 = true
		fmt.Println("b1:", b1) // b1: true
	}
	{
		{
			fmt.Println("############## map 写法1 ##############")
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
			/*
				############## append for map ##############
			*/
			colors["black"] = 23321
			fmt.Println(colors) // map[black:23321 red:49]
			/*
				############## update for map ##############
			*/
			colors["black"] = 99999
			fmt.Println(colors) // map[black:99999 red:49]
			/*
				############## Iterating Over Maps ##############
				k: red v: 49
				k: black v: 99999
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
}

func printMap(inputMap map[string]int) {
	for k, v := range inputMap {
		fmt.Println("k:", k, "v:", v)
	}
}
