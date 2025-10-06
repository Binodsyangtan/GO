package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("example.txt")
	if err != nil {
		//log the error
		panic(err)

	}
	fileInfo, err := file.Stat()
	if err != nil{
		//log the error
		panic(err)
	}

	fmt.Println("file name:", fileInfo.Name())
	fmt.Println("file name:", fileInfo.IsDir())
	fmt.Println("file name:", fileInfo.Size())
	fmt.Println("file name:", fileInfo.Mode())
	fmt.Println("files modified at:", fileInfo.ModTime())


}
