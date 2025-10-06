package main

import (
	
	"fmt"
	"os"
)

func main() {

	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	//log the error
	// 	panic(err)

	// }
	// fileInfo, err := file.Stat()
	// if err != nil{
	// 	//log the error
	// 	panic(err)
	// }

	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file name:", fileInfo.IsDir())
	// fmt.Println("file name:", fileInfo.Size())
	// fmt.Println("file name:", fileInfo.Mode())
	// fmt.Println("files modified at:", fileInfo.ModTime())


	//Read file
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// buffer := make([]byte, 40)
	
	// d,err := file.Read(buffer)
	// if err != nil{
	// 	panic(err)
	// }

	// for i := 0; i < len(buffer); i++{
	// 	println("data",d, string(buffer[i]))
	// }

	//another way not a good way but
	data,err := os.ReadFile("example.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println(string(data))


}
