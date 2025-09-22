package main

import "fmt"
//enumerated types
type OrderStatus int
// type OrderStatus string

const(
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
	// Received OrderStatus = "received"
	// Confirmed            = "confirmed"
)


func changeOrderStatus(status OrderStatus){
	fmt.Println("changing order status to ", status)
}


func main(){
	changeOrderStatus(Confirmed)
	

}