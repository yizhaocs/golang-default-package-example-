package main

import "fmt"

func main()  {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	fmt.Println(mySlice) // [Hi There How Are You]
	updateSlice(mySlice)
	fmt.Println(mySlice) // [Bye There How Are You]
}

// pass by reference的意思是将object以argument pass in到function，可以直接修改value成功
func updateSlice(s []string){
	s[0] = "Bye"
}