package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	// c <- 3
	v, ok := <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
}
