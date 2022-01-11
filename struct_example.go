package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main(){
	{
		yi := person{"Yi", "Zhao"} // 不推荐这种写法，因为change the order of the field会导致出错
		fmt.Println(yi) // {Yi Zhao}
	}

	{
		yi2 := person{firstName: "Yi2", lastName: "Zhao2"} // 推荐这种写法
		fmt.Println(yi2) // {Yi2 Zhao2}
	}

	{
		var yi3 person
		yi3.firstName = "Yi3"
		yi3.lastName = "Zhao3"
		fmt.Println(yi3) // {Yi3 Zhao3}
		fmt.Printf("%+v \n", yi3) // {firstName:Yi3 lastName:Zhao3}
	}
}