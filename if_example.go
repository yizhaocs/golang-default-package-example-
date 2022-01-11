package main

import "fmt"

func main()  {
	fmt.Println("running the program")
	d := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17}
	if len(d) == 17 {
		fmt.Println("len is", len(d)) // len is 17
	} else {
		fmt.Println("len is not", 17) // len is not 17
	}
}

