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

/*
	如果function的argument中的data type定义前有*，则其不算是pointer，只能算是说这个function的这个argument可以接收一个pointer。
	when we see a star in front of a type,
	it means something completely different than when we see a star in front of an actual pointer.
	So in the receiver, we said that pointer to person is a value of type star person.

	Whenever you see a star and then some anything basically any word whatsoever in a place where a type is supposed to be.

	But when we see that star in front of an actual type, that is a type description,
	and it means that this update name function can only be called with the receiver of a pointer to a person.
*/
func (pointerToPerson *person4) updateName(newFirstName string) {
	/*
		So this little parentheses block right here, this thing right here gets turned into the actual yi2
		person that is sitting in memory.
	*/
	(*pointerToPerson).firstName = newFirstName
}

func (p person4) print() {
	fmt.Println("############### func (p person4) print() ###############")
	fmt.Printf("%+v \n", p)
}
