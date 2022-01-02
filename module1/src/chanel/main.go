package main

import "fmt"

func main() {
	go println("a")
	go println("b")
	go println("c")

	ch := make(chan int)

	go func() {
		fmt.Printf("hello from child thread.\n")
		ch <- 20
	}()

	i := <-ch

	fmt.Printf("hello from main.\n")
	fmt.Printf("%d \n", i)

}
