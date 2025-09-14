package main

import "fmt"


// func swap(a, b int) (int,int){
// 	return b, a
	
// }

func main(){
	//arrays fixed sized cannot increase of decrease its sized later
	var marks [5]int = [5]int{90,20,40,30,10}

	fmt.Println("marks:",marks)
	//first start with 0 index means first
	fmt.Println("first student's marks:",marks[0])
	//length of arrya or student
	fmt.Println("total studetns:",len(marks))

	/* A slice is like a dynamic shopping basket in a supermarket.
	You can keep adding items (append).
	You can also remove or rearrange them.
	The basket grows or shrinks as needed. */

	groceries := []string{"milk", "Bread", "eggs"}

	fmt.Println("Groceries:",groceries)

	//add more item
	groceries = append(groceries,"butter", "cheese" )
	fmt.Println("after adding:",groceries)

	//take sub-slice (just some items)
	breakfast := groceries[0:2] //0:2 means 2 items
	fmt.Println("breakfast items:", breakfast)

	for index, item := range groceries {
		fmt.Println(index, ":", item)
	}











	// x, y := 3, 5
	// fmt.Println("before swap:", x,y)

	// //call swap function 
	// x, y = swap(x,y)

	// fmt.Println("after swap:", x, y)

}