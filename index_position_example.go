package main

import "fmt"

func main()  {
	d := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17}
	fmt.Printf("index at 0 is %v \n", d[0]) // index at 0 is 1
	fmt.Printf("index at the last is %v \n", d[len(d) - 1]) // index at the last is 17
}
