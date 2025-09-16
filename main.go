package main

import "fmt"


// func swap(a, b int) (int,int){
// 	return b, a
	
// }

//Define struct
type Student struct {
	Name string
	Age int
	Grade string
}

//product struct
type Product struct{
	ID int
	Name string
	Price float64
	Stock int
}

func main(){

	/* 🔹 What is a Struct?

	A struct is like a blueprint for creating your own data type.

	It groups related data together under one type.

	Similar to objects in JavaScript or classes without methods in Java. */

	//create a studetn
	s1 := Student{Name:"Binod", Age:22 , Grade: "A"}
	s2 :=Student{Name:"sita",Age:21, Grade: "B"}
	fmt.Println("Student 1", s1)
	fmt.Println("Student 2", s2)
	//access fields
	fmt.Println("name of student 1:", s1.Name)

	//create product
	product := Product{ID: 1, Name:"Laptop",Price:75000.50, Stock: 12}
	fmt.Println("Product:", product)
	fmt.Println("Product price:",product.Price)
	

 



	//arrays fixed sized cannot increase of decrease its sized later
	/* var marks [5]int = [5]int{90,20,40,30,10}

	fmt.Println("marks:",marks)
	//first start with 0 index means first
	fmt.Println("first student's marks:",marks[0])
	//length of arrya or student
	fmt.Println("total studetns:",len(marks))
 */
	/* A slice is like a dynamic shopping basket in a supermarket.
	You can keep adding items (append).
	You can also remove or rearrange them.
	The basket grows or shrinks as needed. */

	/* groceries := []string{"milk", "Bread", "eggs"}

	fmt.Println("Groceries:",groceries)

	//add more item
	groceries = append(groceries,"butter", "cheese" )
	fmt.Println("after adding:",groceries)

	//take sub-slice (just some items)
	breakfast := groceries[0:2] //0:2 means 2 items
	fmt.Println("breakfast items:", breakfast)

	for index, item := range groceries {
		fmt.Println(index, ":", item)
	} */

	
	/* 🔹 What is a Map?

	A map is a collection of key-value pairs.

	Think of it like a dictionary in English:

	Word → Definition

	Or a contact list in your phone:

	Name → Phone Number */

	//IN GO
	//map[keyType]ValueTypes

	//create a map
	/* person := map[string]string{
		"name": "Biond",
		"role": "Developer",
	}
	//access values by key
	fmt.Println("Name:",person["name"])
	fmt.Println("Role:",person["role"])

	//contact list
	phonebook :=map[string]string{
		"ram": "98456123",
		"sita": "987456123",
		"binod": "98806455",
	}
	fmt.Println("biond's number:",phonebook["binod"])

	//add new contact
	phonebook["hari"] = "9761245566"

	//delete a contact
	delete(phonebook,"sita")
	
	//loop through map
	fmt.Println("phonebook:")
	for name, number := range phonebook{
		fmt.Println(name, ":", number)
	} */

}