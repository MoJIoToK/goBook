package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)

}

// printer - function for printing output number
func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

// squarer - function for square number.
func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

// counter - function for generation number.
func counter(out chan<- int) {
	for x := 0; x < 4; x++ {
		out <- x
		time.Sleep(1 * time.Second)
	}
	close(out)
}
