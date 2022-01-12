package main

import "fmt"

func main()  {
	myInt := 3
	fmt.Println(myInt) // 3
	updateInt(myInt)
	fmt.Println(myInt) // 3
	myIntPointer := &myInt
	updateInt2(myIntPointer)
	fmt.Println(myInt) // 11
}

// update failed
func updateInt(passInt int){
	passInt = 9
}

// update successfully
func updateInt2(passInt *int){
	(*passInt) = 11
}