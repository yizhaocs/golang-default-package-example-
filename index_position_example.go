package main

import "fmt"

func main()  {
	slice := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17}
	fmt.Printf("index at 0 is %v \n", slice[0]) // index at 0 is 1
	fmt.Printf("index at the last is %v \n", slice[len(slice) - 1]) // index at the last is 17
	fmt.Printf("从index 0到index2的elements %v \n",slice[0:2]) // 从index 0到index2的elements [1 2]
	fmt.Printf("从index 3开始 %v \n",slice[3:]) // 从index 3开始 [4 5 6 7 8 9 10 11 12 13 14 15 16 17]
	fmt.Printf("最后3个elements %v \n",slice[:3]) // 最后3个elements [1 2 3]
}
