package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main()  {
	error_example("write_to_filename.txt")
}

func error_example(filename string)  {
	bytes, error := ioutil.ReadFile(filename)
	if error != nil { // 如果出错为nil则print出错误并且退出程序
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	fmt.Println("file found with content:", string(bytes))
}

