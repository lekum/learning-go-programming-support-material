package main

import "fmt"

func main() {
	ch := make(chan int, 4)
	ch <- 2
	ch <- 3
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}
