package main

import "fmt"

type personhaha struct {
	firstName string
	lastName  string
}

func main() {
	{
		/*
			print basic
		*/
		fmt.Println("running the program")  // running the program
		fmt.Println("A", 1, "B", 2, "C", 3) // A 1 B 2 C 3
	}
	{
		/*
			print slice
		*/
		d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
		fmt.Println("len is", len(d))          // len is 17
		fmt.Println("len is", len(d), "d:", d) // len is 17 d: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17]
		fmt.Printf("len is %v \n", len(d))     // len is 17
	}
	{
		/*
			print出struct里的所有fields用"%+v"
		*/
		var yi3 personhaha
		yi3.firstName = "Yi3"
		yi3.lastName = "Zhao3"
		fmt.Println(yi3)          // {Yi3 Zhao3}
		fmt.Printf("%+v \n", yi3) // {firstName:Yi3 lastName:Zhao3}
	}

	fmt.Println("end the program")
}
