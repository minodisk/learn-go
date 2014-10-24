package main

import (
	"log"

	"../tree"
)

func Walk(t *tree.Tree, c chan int) {
	walk(t, c)
	close(c)
}

func walk(t *tree.Tree, c chan int) {
	if t != nil {
		walk(t.Left, c)
		c <- t.Value
		walk(t.Right, c)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := range c1 {
		if i != <-c2 {
			return false
		}
	}
	return true
}

func main() {
	c := make(chan int)
	log.Println(tree.New(1).Value)
	go Walk(tree.New(1), c)
	for v := range c {
		log.Println(v)
	}
}
