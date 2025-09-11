package main

import "fmt" //fmt stand for format

func main() {
	var days int

	fmt.Println("enter your marks:")
	fmt.Scanln(&days)  //taking input

	switch days {
	case 1:
		fmt.Println("today is sunday")

	case 2:
		fmt.Println("today is monday")

	case 3:
		fmt.Println("today is tuesday")

	case 4:
		fmt.Println("today is wednesday")
	case 5:
		fmt.Println("today is thrusday:")

	case 6:
		fmt.Println("today is friday")

	case 7:
		fmt.Println("today is saturday")

	default:
		fmt.Println("invalid day")
		
	}


}
