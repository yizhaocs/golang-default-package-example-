package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	error_example("xyz.txt")
}

func error_example(filename string)  {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("file found with content:", string(bytes))
}
