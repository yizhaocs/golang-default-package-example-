package main

import "fmt"

func main()  {
	greating := "Hi there!"
	slice_byte := []byte(greating) // string -> slice of bytes
	fmt.Println("slice_byte:", slice_byte) // slice_byte: [72 105 32 116 104 101 114 101 33]
	string_data_type := string(slice_byte) // slice of bytes -> string
	fmt.Println("string_data_type:", string_data_type) // string_data_type: Hi there!
}
