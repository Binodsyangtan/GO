package main

import "fmt"
/* 🔹 What is Generics?

Normally in Go, if you want a function that works with multiple types (like int, float64, string), you have to duplicate code or use interface{} (which loses type safety).

Generics let you write functions and data structures that work with different types while keeping type safety.

👉 In short: Generics = Reusable code with type parameters.

🔹 Real-World Analogy

Think of a shopping bag 🛍️.

You don’t want a bag just for apples, and another just for bananas.

Instead, you want one bag that can hold any item, but still ensures they are valid items.

Generics give us this "flexible bag" in Go. */



//we have to write same line of code again and again with just little change 
//with generic we can write code with reusabe (DONOT REPEAT CODE)

//any is not good to use sometime
// we can used int and string limited as well or interface{} like this as well
//used bool as well so we can pass bool too

//comparable ( we can used this instead of any  or int | string | bool)
func printSlice [T int | string | bool] (items []T){
	for _, item :=range items{
		fmt.Println(item)
	}
}

//string type code 
// func printStringSlice (items []string){
// 	for _, item := range items{
// 		fmt.Println(item)
// 	}
// }

//struct
//LIFO
type stack[T any] struct{
	elements []T
}

func main(){
	myStack := stack[string] {
		elements : []string{"golang"},
	}
	fmt.Println(myStack)




	nums := []int{1,2,3}
	names := []string{"go","js","java"}
	vals := []bool{true,false,true}
	// printStringSlice(StringSlice)
	printSlice(names) //detect it is string without writing string code
	printSlice(nums)  //now int type
	printSlice(vals) // we didnot use bool so we connot used bool value for that we have to pass bool as well

}