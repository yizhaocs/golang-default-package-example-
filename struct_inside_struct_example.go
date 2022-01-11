package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}


type person2 struct {
	firstName string
	lastName string
	contact contactInfo
}

func main(){
	{
		yi2 := person2{
			firstName: "Yi2",
			lastName: "Zhao2",
			contact: contactInfo{
				email: "yi@gmail.com",
				zipCode: 94000,
			},
		} // 推荐这种写法
		fmt.Printf("%+v \n", yi2) // {firstName:Yi2 lastName:Zhao2 contact:{email:yi@gmail.com zipCode:94000}}
	}
}