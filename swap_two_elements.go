package main

import (
	"fmt"
)

func main()  {
	a, b := "second", "first"
	fmt.Println(a, b) // Prints "second first"
	b, a = a, b
	fmt.Println(a, b) // Prints "first second"
}

