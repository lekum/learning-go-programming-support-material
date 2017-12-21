package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int
	fmt.Println("A square is what?")
	fmt.Println("Enter 1=quadrilateral, 2=rectagonal:")
	n, err := fmt.Fscanf(os.Stdin, "%d", &choice)
	if err != nil || n != 1 {
		fmt.Println("Follow directions!:", err)
		os.Exit(1)
	}
	if choice == 1 {
		fmt.Println("You are correct!")
	} else {
		fmt.Println("Wrong. Google it.")
	}
}
