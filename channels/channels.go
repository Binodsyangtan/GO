package main

import (
	"fmt"
)

//send
// func processNum(numChan chan int) {

// 	for num :=range numChan{
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second * 2)

// 	}

// }



//receive
// func sum(result chan int, num1 int, num2 int) {
// 	numResult := num1 + num2
// 	result <- numResult

// }

func main() {

	//receive
	// result := make(chan int)
	// go sum(result, 4, 5)

	// res := <-result //blocking                    
	// fmt.Println("sum of two int", res)

	//send
	// numChan := make(chan int)

	// go processNum(numChan)

	// for{
	// 	numChan <- rand.Intn(100) //random number 0-100 generate garxa

	// }

	// time.Sleep(time.Second * 2)

	// messageChan := make(chan string) //create channel

	// messageChan <- "ping" //send data   //blocking

	// msg := <-messageChan //Receive data

	// fmt.Println(msg)
	// // fatal error: all goroutines are asleep - deadlock!  (create channel send msg and receive msg matra gare yesto error dinxa)

}
