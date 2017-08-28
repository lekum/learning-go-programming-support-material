package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error parsing arg:", r)
			os.Exit(1)
		}
	}()
	arg := os.Args[1]
	fmt.Println(arg)
}
