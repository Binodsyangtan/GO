package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() //add one gareko thiyo usley hatauxa feri one lai
	fmt.Println("doing task", id)
}

func main() {

	//declare waitgroup
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1) //add one
		go task(i, &wg)

	}
	// time.Sleep(time.Second * 2)  //for this we used waitgroup
	
	//now wait garni 
	wg.Wait()
}
