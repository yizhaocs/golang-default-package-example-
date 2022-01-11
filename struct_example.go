package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main(){
	yi := person{"Yi", "Zhao"} // 不推荐这种写法，因为change the order of the field会导致出错
	yi2 := person{firstName: "Yi2", lastName: "Zhao2"} // 推荐这种写法

	fmt.Println(yi) // {Yi Zhao}
	fmt.Println(yi2) // {Yi2 Zhao2}
}