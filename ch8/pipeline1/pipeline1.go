package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//Generate number
	go func() {
		for x := 0; ; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
	}()

	//Square number
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	//Print output number
	for {
		fmt.Println(<-squares)
	}

}
