package main

import (
	"fmt"
	"unicode"
)

func main() {
	{
		fmt.Println("******************************************************************** ")
		fmt.Println("******************************* IsLetter和IsNumber *******************************")
		fmt.Println("******************************************************************** ")
		fmt.Println("***************** IsLetter ***************** ")
		fmt.Println(unicode.IsLetter('1')) // false
		fmt.Println(unicode.IsLetter('a')) // true

		fmt.Println("***************** IsNumber ***************** ")
		fmt.Println(unicode.IsNumber('1')) // true
		fmt.Println(unicode.IsNumber('a')) // false
	}
}