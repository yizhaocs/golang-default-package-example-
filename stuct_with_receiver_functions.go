package main

import "fmt"

type contactInfo2 struct {
	email   string
	zipCode int
}

type person4 struct {
	firstName string
	lastName  string
	contact   contactInfo2
}

func main() {
	yi2 := person4{
		firstName: "Yi2",
		lastName:  "Zhao2",
		contact: contactInfo2{
			email:   "yi@gmail.com",
			zipCode: 94000,
		},
	} // 推荐这种写法
	{
		fmt.Printf("%+v \n", yi2) // {firstName:Yi2 lastName:Zhao2 contact:{email:yi@gmail.com zipCode:94000}}
	}
	{
		/*
			############### func (p person4) print() ###############
			{firstName:Yi2 lastName:Zhao2 contact:{email:yi@gmail.com zipCode:94000}}
		*/
		yi2.print()
	}

	{
		yi2Pointer := &yi2 // 因为golang是一种pass by value语言，所以要用pointer将原内存地址赋予到pointer
		yi2Pointer.updateName("abcd")
		/*
			############### func (p person4) print() ###############
			{firstName:abcd lastName:Zhao2 contact:{email:yi@gmail.com zipCode:94000}}
		*/
		yi2.print()
	}
}

func (pointerToPerson *person4) updateName(newFirstName string) {
	pointerToPerson.firstName = newFirstName
}

func (p person4) print() {
	fmt.Println("############### func (p person4) print() ###############")
	fmt.Printf("%+v \n", p)
}
