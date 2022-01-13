package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	{
		fmt.Println("############## 初始化map用make ##############")
		var colors2 map[string]int
		fmt.Println(colors2) // map[]
		colors2 = make(map[string]int)
		colors2["pink"] = 213
		fmt.Println(colors2) // map[pink:213]
	}

	{
		response, error := http.Get("http://google.com")
		if error != nil {
			fmt.Println("Error:", error)
			os.Exit(1)
		}

		/*
			So this make function right here is a built-In function in the language that takes a type of a slice.
			And then as a second argument,
			this is the number of elements or empty spaces that we want to slice to be initialized with.
		*/
		bs := make([]byte, 99999)
		response.Body.Read(bs)
	}
}
