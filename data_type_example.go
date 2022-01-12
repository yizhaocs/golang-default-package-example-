package main

import "fmt"

func main()  {
	{
		var s1 string
		s1 = "s1"
		fmt.Println("string:", s1) // string: s1
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
}
