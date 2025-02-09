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
		for x := 0; x < 4; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	//Square number
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	//Print output number
	for x := range squares {
		fmt.Println(x)
	}

}
