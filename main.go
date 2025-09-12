package main

import "fmt"


func swap(a, b int) (int,int){
	return b, a
	
}

func main(){
	x, y := 3, 5
	fmt.Println("before swap:", x,y)

	//call swap function 
	x, y = swap(x,y)

	fmt.Println("after swap:", x, y)

}