package main

import (
	"fmt"
	"os"
	"strconv"
)

func dbl(p *float64) {
	*p *= 2
}

func main() {
	num, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("Error parsing float: %v\n", err)
		os.Exit(1)
	}
	dbl(&num)
	fmt.Println(num)
}
