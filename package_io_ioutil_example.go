package main

import (
	"io/ioutil"
)

func main()  {
	saveToFile("write_to_filename.txt")
}

func  saveToFile(filename string) error {
	ex1 := "Ace of Spades 1"
	// 0666 means anyone can read and write this file
	return ioutil.WriteFile(filename, []byte(ex1), 0666)
}
